# GoFace

A university project in Golang for face recognition.

## How to run?

In order to run the project, execute the command `go run ./cmd/GoFace`, while being in the main project directory.

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
