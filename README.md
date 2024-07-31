# gotth-components

A component library powered by golang+templ+tailwindcss+htmx

---

### Local Setup

1. Clone the repository

```bash
git clone https://github.com/m3rashid/gotth-components.git
```

2. Install air (for hot reloading)

```bash
go install github.com/air-verse/air@latest
```

3. Install the templ cli

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

4. Setup tailwindcss

   1. Download the binary (depending on your operating system) from `https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.7`
   2. Extract the binary and move it to the project root
   3. rename it to `tailwindcss`
   4. Make it executable `chmod +x ./tailwindcss`
   5. Run `make tailwind-watch` to generate style.css file
   6. Run `make tailwind-build` to generate style.min.css file for production

5. Install the dependencies `go mod download`
6. Run the project `make dev`

### Readings

Please checkout the [Docs](./docs/) for [Project Structure](./docs/structure.md) and [Project Decisions](./docs/decisions.md)
