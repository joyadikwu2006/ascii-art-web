# ascii-art-web

## Description
A web-based GUI application that converts text into ASCII art using different banner styles.
Built in Go, it wraps the ascii-art CLI project in an HTTP server with a browser-based interface.
Users can type any text, choose a banner style, and see the generated ASCII art rendered on the page.

## Authors
- Joy Adikwu

## Usage

### Requirements
- Go 1.18 or higher

### How to run
1. Clone the repository
2. Navigate to the project directory
cd ascii-art-web
3. Run the server
go run main.go
4. Open your browser and visit
http://localhost:8080

## Implementation Details

### Algorithm
1. The server starts and registers two HTTP routes: `GET /` and `POST /ascii-art`
2. `GET /` serves the main HTML page with a text input, banner style selector, and submit button
3. When the user submits the form, the browser sends a `POST /ascii-art` request with the text and chosen banner style as form values
4. The server reads the text and banner style from the request using `r.FormValue()`
5. It loads the corresponding banner file from the `banners/` directory using `LoadBanner()`, which parses the `.txt` file into a map of rune → 8 rows of ASCII art
6. It passes the text and banner map into `GenerateArt()`, which splits the input on `\n`, renders each line row by row using the banner map, and joins the rows into a final string
7. The result, along with the original text and banner style, is passed into the HTML template so the form stays pre-filled and the ASCII art is displayed in a `<pre>` block
8. All errors return appropriate HTTP status codes: 400 for bad input, 404 for missing templates or banners, 500 for unhandled server errors

## Known Limitations

**Partial-write errors during template rendering.** The `/ascii-art` handler
calls `tmpl.Execute(w, ...)` to render the result directly onto the response
writer. If `Execute` fails partway through — for example, due to a template
error after some HTML has already been written — the response headers and
part of the body have already been sent to the client with a `200 OK` status.
The subsequent call to `http.Error` in that case cannot change the status
code or replace the already-sent body; it can only append error text to the
end of the partial page. A more robust fix would render into a
`bytes.Buffer` first and only write to `w` (with the correct status code)
once rendering succeeds in full — left as a future improvement.
