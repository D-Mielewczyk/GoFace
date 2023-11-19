# GoFace

A university project in Golang for face recognition.

## Running GoFace

### Using Go Command

To run the GoFace project with Go installed, open a command prompt or terminal and navigate to the main project directory. Then execute the following command:

```powershell
go run ./cmd/GoFace
```

This method requires Go to be installed on your system.

### Using the Executable

If you prefer not to install Go or are working on a system without Go, you can use the GoFace.exe executable for ease. Simply replace the command go run ./cmd/GoFace with GoFace.exe in all usage instructions. For instance:

```powershell
GoFace.exe <arguments>
```

This executable allows you to run the GoFace project directly, without needing Go installed.

## CLI Usage

The GoFace CLI provides a simple interface to run face detection on images. You can specify an individual image or process all images within a directory. Additionally, you have the option to draw either circles or rectangles around detected faces.

### Commands

1. **Detect Face in a Single Image:**
    To detect faces in a single image, provide the image file name (relative to the `images` directory) as an argument.

    ```powershell
    go run ./cmd/GoFace <image_file_name>
    ```

    For example, if you have an image named `ja.jpg` in the `images` directory:

    ```powershell
    go run ./cmd/GoFace ja.jpg
    ```

2. **Detect Faces in All Images in a Directory:**

    To run face detection on all images in the `images` directory, use the `all` argument.

    ```powershell
    go run ./cmd/GoFace all
    ```

3. **Option to Draw Circles or Rectangles:**

By default, rectangles are drawn around detected faces. You can choose to draw circles instead by adding `circle` as the last argument.

- For a single image:

  ```powershell
  go run ./cmd/GoFace <image_file_name> circle
  ```

- For all images:

  ```powershell
  go run ./cmd/GoFace all circle
  ```

### Output

The processed images will be saved in the `output` directory with the same file names as the input images.
