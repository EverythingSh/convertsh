package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EverythingSh/convertsh/pkg/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// FileType represents a broad file type category
type FileType int

const (
	TypeImages FileType = iota
	TypeVideos
	TypeAudio
)

var fileTypeNames = []string{"Images (supported)", "Videos (unsupported)", "Audio (unsupported)"}

// File represents a file or directory in the file system
type File struct {
	Name     string
	Path     string
	IsDir    bool
	Selected bool
}

// TUIStage represents the current stage of the TUI
type TUIStage int

const (
	StageSelectType TUIStage = iota
	StageBrowseFiles
	StageSelectFormat
	StageDone
)

// Model represents the application state
type Model struct {
	Stage          TUIStage
	FileType       FileType
	Files          []File
	Cursor         int
	CurrentDir     string
	Width          int
	Height         int
	Message        string
	SelectedFiles  map[string]File
	Formats        []types.ImageRasterFormat // Use types.FileFormat from pkg/types
	FormatCursor   int
	SelectedFormat *types.ImageRasterFormat
}

// Styles (same as before)
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7aa2f7")).
			Bold(true).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9ece6a")).
			Bold(true)

	directoryStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e0af68"))

	fileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#c0caf5"))

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#bb9af7"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7aa2f7")).
			MarginTop(1).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#565f89")).
			Align(lipgloss.Left)
)

// InitialModel creates a new reusable TUI model
func InitialModel() Model {
	return Model{
		Stage:         StageSelectType,
		Cursor:        0,
		Width:         50,
		Height:        20,
		SelectedFiles: make(map[string]File),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch m.Stage {
		case StageSelectType:
			switch msg.String() {
			case "up", "k":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down", "j":
				if m.Cursor < len(fileTypeNames)-1 {
					m.Cursor++
				}
			case "enter":
				m.FileType = FileType(m.Cursor)
				if m.FileType != TypeImages {
					m.Message = "Only images are supported for now."
					return m, nil
				}
				// Move to file browser for images
				m.Stage = StageBrowseFiles
				m.Cursor = 0
				m.CurrentDir, _ = os.Getwd()
				files, _ := getFilesInDirFiltered(m.CurrentDir, m.FileType)
				m.Files = files
				m.Message = ""
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case StageBrowseFiles:
			switch msg.String() {
			case "up", "k":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down", "j":
				if m.Cursor < len(m.Files)-1 {
					m.Cursor++
				}
			case " ", "s":
				selectedFile := m.Files[m.Cursor]
				if selectedFile.IsDir {
					newDir := selectedFile.Path
					files, err := getFilesInDirFiltered(newDir, m.FileType)
					if err != nil {
						m.Message = fmt.Sprintf("Error: %v", err)
						return m, nil
					}
					m.Files = files
					m.CurrentDir = newDir
					m.Cursor = 1
					m.Message = ""
				} else {
					m.Files[m.Cursor].Selected = !m.Files[m.Cursor].Selected
					if m.Files[m.Cursor].Selected {
						m.SelectedFiles[selectedFile.Path] = selectedFile
					} else {
						delete(m.SelectedFiles, selectedFile.Path)
					}
				}
			case "backspace", "h", "left":
				if m.CurrentDir != "/" {
					parentDir := filepath.Dir(m.CurrentDir)
					files, err := getFilesInDirFiltered(parentDir, m.FileType)
					if err != nil {
						m.Message = fmt.Sprintf("Error: %v", err)
						return m, nil
					}
					m.Files = files
					m.CurrentDir = parentDir
					m.Cursor = 0
					m.Message = ""
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter":
				selectedFile := m.Files[m.Cursor]
				if selectedFile.IsDir {
					newDir := selectedFile.Path
					files, err := getFilesInDirFiltered(newDir, m.FileType)
					if err != nil {
						m.Message = fmt.Sprintf("Error: %v", err)
						return m, nil
					}
					m.Files = files
					m.CurrentDir = newDir
					m.Cursor = 0
					m.Message = ""
				} else {
					m.SelectedFiles[selectedFile.Path] = selectedFile
				}
				// Proceed to format selection if at least one file is selected
				if len(m.SelectedFiles) > 0 {
					m.Stage = StageSelectFormat
					m.Formats = types.RasterFormats // Use formats from pkg/types
					m.FormatCursor = 0
				}
			}

		case StageSelectFormat:
			switch msg.String() {
			case "up", "k":
				if m.FormatCursor > 0 {
					m.FormatCursor--
				}
			case "down", "j":
				if m.FormatCursor < len(m.Formats)-1 {
					m.FormatCursor++
				}
			case "enter":
				m.SelectedFormat = &m.Formats[m.FormatCursor]
				m.Stage = StageDone
			case "q", "ctrl+c":
				return m, tea.Quit
			case "backspace":
				m.Stage = StageBrowseFiles
			}
		case StageDone:
			switch msg.String() {
			case "q", "ctrl+c", "enter":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.Stage {
	case StageSelectType:
		s := titleStyle.Render("Select File Type") + "\n\n"
		for i, name := range fileTypeNames {
			cursor := " "
			if m.Cursor == i {
				cursor = cursorStyle.Render("•")
			}
			s += fmt.Sprintf(" %s %s\n", cursor, name)
		}
		s += "\n" + helpStyle.Render("↑/↓: navigate • enter: select • q: quit")
		if m.Message != "" {
			s += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#f7768e")).Render(m.Message)
		}
		return s

	case StageBrowseFiles:
		s := titleStyle.Render("Select Images") + "\n"
		s += statusStyle.Render(m.CurrentDir) + "\n\n"
		visibleItems := m.Height - 8
		if visibleItems < 1 {
			visibleItems = 10
		}
		start := 0
		if len(m.Files) > visibleItems {
			middle := visibleItems / 2
			if m.Cursor > middle {
				start = m.Cursor - middle
			}
			if start+visibleItems > len(m.Files) {
				start = len(m.Files) - visibleItems
			}
			if start < 0 {
				start = 0
			}
		}
		end := start + visibleItems
		if end > len(m.Files) {
			end = len(m.Files)
		}
		for i := start; i < end; i++ {
			file := m.Files[i]
			cursor := " "
			if m.Cursor == i {
				cursor = cursorStyle.Render("•")
			}
			fileLabel := file.Name
			if file.IsDir {
				fileLabel = directoryStyle.Render(fileLabel + "/")
			} else {
				if file.Selected {
					fileLabel = selectedStyle.Render("✓ " + fileLabel)
				} else {
					fileLabel = fileStyle.Render(fileLabel)
				}
			}
			s += fmt.Sprintf(" %s %s\n", cursor, fileLabel)
		}
		if len(m.Files) > visibleItems {
			s += helpStyle.Render(fmt.Sprintf("\n%d-%d of %d", start+1, end, len(m.Files)))
		}
		if m.Message != "" {
			s += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#f7768e")).Render(m.Message)
		}
		s += "\n\n" + helpStyle.Render("↑/↓: navigate • space: select • enter: open • tab: next • backspace: back • q: quit")
		return s

	case StageSelectFormat:
		s := titleStyle.Render("Select Output Format") + "\n\n"
		for i, f := range m.Formats {
			cursor := " "
			if m.FormatCursor == i {
				cursor = cursorStyle.Render("•")
			}
			s += fmt.Sprintf(" %s %s\n", cursor, f)
		}
		s += "\n" + helpStyle.Render("↑/↓: navigate • enter: select • backspace: back • q: quit")
		return s

	case StageDone:
		s := titleStyle.Render("Selection Complete!") + "\n\n"
		s += "Selected files:\n"
		files, targetFormat := m.GetSelectedFiles()
		for _, f := range files {
			s += fmt.Sprintf("  %s\n", f.Path)
		}
		s += fmt.Sprintf("\nTarget format: %s\n", *targetFormat)
		s += "\nPress enter or q to quit."
		return s

	default:
		return "Unknown stage"
	}
}

// getFilesInDirFiltered returns files/directories filtered by file type
func getFilesInDirFiltered(dir string, fileType FileType) ([]File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []File
	if dir != "/" {
		files = append(files, File{
			Name:  "..",
			Path:  filepath.Dir(dir),
			IsDir: true,
		})
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") && entry.Name() != ".." {
			continue
		}
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			files = append(files, File{
				Name:  entry.Name(),
				Path:  fullPath,
				IsDir: true,
			})
		} else if fileType == TypeImages && isImageFile(entry.Name()) {
			files = append(files, File{
				Name:  entry.Name(),
				Path:  fullPath,
				IsDir: false,
			})
		}
	}
	return files, nil
}

// isImageFile checks if a file is an image by extension
func isImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp":
		return true
	default:
		return false
	}
}

// GetSelectedFiles returns all selected files
func (m Model) GetSelectedFiles() ([]File, *types.ImageRasterFormat) {
	var selectedFiles []File
	for _, file := range m.SelectedFiles {
		selectedFiles = append(selectedFiles, file)
	}
	return selectedFiles, m.SelectedFormat
}
