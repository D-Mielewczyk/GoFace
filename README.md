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

The GoFace CLI provides a simple interface to run face detection on images. You can specify an individual image or process all images within a directory. Additionally, you have the option to draw either circles or rectangles around detected faces, specify the output directory, get help information, and measure performance.

### Commands

1. **Detect Face in a Single Image:**
    To detect faces in a single image, provide the image file name (relative to the `images` directory) as an argument.

    ```powershell
    go run ./cmd/GoFace -o outputdir <image_file_name>
    ```

    For example, if you have an image named `ja.jpg` in the `images` directory:

    ```powershell
    go run ./cmd/GoFace -o outputdir ja.jpg
    ```

2. **Detect Faces in All Images in a Directory:**

    To run face detection on all images in the `images` directory, use the `all` argument.

    ```powershell
    go run ./cmd/GoFace -o outputdir all
    ```

3. **Option to Draw Circles or Rectangles:**

    By default, rectangles are drawn around detected faces. You can choose to draw circles instead by using the `-c` flag.

    - For a single image:

      ```powershell
      go run ./cmd/GoFace -c -o outputdir <image_file_name>
      ```

    - For all images:

      ```powershell
      go run ./cmd/GoFace -c -o outputdir all
      ```

4. **Display Help Information:**

    To display help information about the available commands and flags, use the `-h` flag.

    ```powershell
    go run ./cmd/GoFace -h
    ```

5. **Measure Performance:**

    To measure and display the performance of the face detection process, use the `-p` flag.

    - For a single image:

      ```powershell
      go run ./cmd/GoFace -p -o outputdir <image_file_name>
      ```

    - For all images:

      ```powershell
      go run ./cmd/GoFace -p -o outputdir all
      ```

### Output

By default the processed images will be saved in the `output` directory with the same file names as the input images.

## Running Tests

To execute the tests for this project, use the following command in the terminal or command prompt:

```powershell
go test -v ./tests    
```

his command will run all the tests located in the tests directory. The -v flag provides verbose output, showing detailed information about each test case.

If you prefer a more concise output, indicating only whether the tests passed or failed, you can omit the -v flag:

```powershell
go test ./tests
```

## Acknowledgments and References

### Pigo - Pure Go Face Detection Library

This project utilizes [Pigo](https://github.com/esimov/pigo), a pure Go library for face detection, as a key component in the face detection process. Pigo provides an efficient and easy-to-use way to detect faces in images without any external dependencies.

For more information about Pigo and its usage, visit the [Pigo GitHub repository](https://github.com/esimov/pigo/tree/master).
