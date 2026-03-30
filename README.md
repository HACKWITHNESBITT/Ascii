      ____    _____                      _____    _____   _____   _____  
     / __ \  |  __ \            /\      / ____|  / ____| |_   _| |_   _| 
    | |  | | | |__) |          /  \    | (___   | |        | |     | |   
    | |  | | |  _  /          / /\ \    \___ \  | |        | |     | |   
    | |__| | | | \ \         / ____ \   ____) | | |____   _| |_   _| |_  
     \___\_\ |_|  \_\       /_/    \_\ |_____/   \_____| |_____| |_____| 
                                                                     




## Project Overview

      The QR ASCII Web App is a web application that allows users to generate QR codes from any input text or URL and view them in two formats:
      Image format – standard QR code image (PNG)
      ASCII art format – QR code represented using ASCII characters, like ██ and spaces

This project is built using Go (Golang) for the backend and standard HTML, CSS, and JavaScript for the frontend.

## Key Features

      QR Code Generation: Convert text, URLs, or data into QR codes using the go-qrcode library.
      ASCII Art Rendering: QR codes can be displayed as ASCII art in the browser or terminal.
      Download Options: Users can download QR codes as:
      PNG image
      ASCII text file
      API Support: Generate QR codes programmatically through an API endpoint (/api/qr).
      Interactive Frontend: Supports drag-and-drop input, dark/light mode toggle, and responsive design.
      How It Works
      Frontend Input:
      The user enters text or a URL in the web interface.
      Backend Processing (Go):
      Go receives the input via HTTP handlers (api/qr.go).
      Using go-qrcode, the input is converted into a QR code matrix (a 2D array of black/white modules).
      The backend can generate:
      A PNG image of the QR code
      An ASCII art representation, mapping true cells to ██ and false cells to spaces
      Display & Download:
      The QR code is displayed in the browser in both formats.
      Users can download the generated QR code as a PNG image or ASCII file.

## Folder Structure

      qr-ascii-web/
      ├── main.go             # Entry point: sets up server and routes
      ├── go.mod              # Go module dependencies
      ├── api/
      │   └── qr.go           # QR code HTTP handlers and API logic
      ├── utils/
      │   └── ascii.go        # Converts QR code matrix to ASCII art
      ├── templates/
      │   └── index.html      # Frontend HTML template
      ├── static/
      │   ├── style.css       # CSS for styling
      │   └── script.js       # JavaScript for user interaction
      └── README.md

## Technologies Used

      Go (Golang): Backend logic, QR code generation, and HTTP server
      go-qrcode: Library to generate QR code matrices
      HTML, CSS, JS: Frontend UI and interactivity
      ASCII Art: For visual representation in text format

Use Cases

      Share QR codes as ASCII in emails or chat where images are not supported.
      Quickly generate QR codes for URLs, Wi-Fi credentials, or text.
      Educational purpose: visualize how QR codes are structured.
