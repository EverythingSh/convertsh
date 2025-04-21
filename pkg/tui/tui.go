package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// File represents a file or directory in the file system
type File struct {
	Name     string
	Path     string
	IsDir    bool
	Selected bool
}

// Model represents the application state
type Model struct {
	Files         []File
	Cursor        int
	CurrentDir    string
	Width         int
	Height        int
	Message       string
	SelectedFiles map[string]File
}

// Styles
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

func InitialModel() Model {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return Model{}
	}

	files, err := getFilesInDir(currentDir)
	if err != nil {
		return Model{Message: fmt.Sprintf("Error: %v", err)}
	}

	return Model{
		Files:         files,
		Cursor:        0,
		CurrentDir:    currentDir,
		Width:         50,
		Height:        20,
		SelectedFiles: make(map[string]File),
	}
}

func getFilesInDir(dir string) ([]File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []File
	// Add parent directory entry
	if dir != "/" {
		files = append(files, File{
			Name:  "..",
			Path:  filepath.Dir(dir),
			IsDir: true,
		})
	}

	for _, entry := range entries {
		// Skip hidden files
		if strings.HasPrefix(entry.Name(), ".") && entry.Name() != ".." {
			continue
		}

		files = append(files, File{
			Name:  entry.Name(),
			Path:  filepath.Join(dir, entry.Name()),
			IsDir: entry.IsDir(),
		})
	}
	return files, nil
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
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Files)-1 {
				m.Cursor++
			}

		case "home":
			m.Cursor = 0

		case "end":
			m.Cursor = len(m.Files) - 1

		case "enter":
			selectedFile := m.Files[m.Cursor]
			if selectedFile.IsDir {
				// Navigate into the directory
				newDir := selectedFile.Path
				files, err := getFilesInDir(newDir)
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

		case " ":
			if !m.Files[m.Cursor].IsDir {
				m.Files[m.Cursor].Selected = !m.Files[m.Cursor].Selected
			}

		case "backspace", "h", "left":
			// Go up one directory
			if m.CurrentDir != "/" {
				parentDir := filepath.Dir(m.CurrentDir)
				files, err := getFilesInDir(parentDir)
				if err != nil {
					m.Message = fmt.Sprintf("Error: %v", err)
					return m, nil
				}
				selectedFiles := m.GetSelectedFiles()
				for _, file := range selectedFiles {
					for i, f := range files {
						if file.Path == f.Path {
							files[i].Selected = true
						}
					}
				}
				m.Files = files
				m.CurrentDir = parentDir
				m.Cursor = 0
				m.Message = ""
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if len(m.Files) == 0 {
		return titleStyle.Render("Empty directory") + "\n" +
			helpStyle.Render("\nPress 'backspace' to go up, 'q' to quit")
	}

	// Title bar - simple and clean
	s := titleStyle.Render("File Browser")

	// Current directory - clean with minimal decoration
	s += fmt.Sprintf("\n%s\n\n", statusStyle.Render(m.CurrentDir))

	// Visible height calculation
	visibleItems := m.Height - 8 // Reduced space for UI elements
	if visibleItems < 1 {
		visibleItems = 10
	}

	// Calculate pagination
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

	// Files list with minimal styling
	for i := start; i < end; i++ {
		file := m.Files[i]
		cursor := " "
		if m.Cursor == i {
			cursor = "•" // Simple bullet point cursor
			cursor = cursorStyle.Render(cursor)
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

	// Simple pagination indicator
	if len(m.Files) > visibleItems {
		s += helpStyle.Render(
			fmt.Sprintf("\n%d-%d of %d", start+1, end, len(m.Files)))
	}

	// Error message - simple inline notification
	if m.Message != "" {
		s += "\n" + lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f7768e")).
			Render("Error: "+m.Message)
	}

	// Help text - simplified and minimal
	s += "\n\n" + helpStyle.Render("↑/↓: navigate • space: select • enter: open • backspace: back • q: quit")

	return s
}

// GetSelectedFiles returns all selected files
func (m Model) GetSelectedFiles() []File {
	var selectedFiles []File
	for _, file := range m.SelectedFiles {
		selectedFiles = append(selectedFiles, file)
	}
	return selectedFiles
}
