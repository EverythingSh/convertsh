# ConvertSh
Convert anything you like : )
**The project is currently under development🏗️**

## TO-DO

- Support for bulk conversion of files.
- Support for zip/unzip and file compression.
- Support for almost all image, video, audio, and document types.
- Support for a beautiful TUI for navigating directories and selecting files for conversion.
- Support for accessing Google Drive workspace for importing files and conversion.
- Easy installation of the tool using different package managers.
- Cross-platform support.
- Support of AI for enhancing audio containing background noises, enhancing images like removing background from objects, etc.

## Installing Dependencies
I've used go-vips(a wrapper around class for libvips). Dependencies that are needed.

1) If on OSX then set:
```
export CGO_CFLAGS_ALLOW="-Xpreprocessor"
```

2) Install to packages namely `pkg-config` and `vips`.

## Building the Project

To build the project, use the following command:

```sh
go build -o build/con cmd/main.go
```

## Running the Project

To run the project, use the following command:

```sh
./build/con input.jpg
```

Currently, the app expects a JPEG file to be passed as an argument and only converts it to a PNG file.
