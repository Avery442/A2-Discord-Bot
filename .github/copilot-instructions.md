# A2 Discord Bot

A2 Discord Bot is a Go application that connects to Discord and provides information about A2 game server fleets. The bot responds to the command "howmanyspacemonke" with a formatted table showing station names, versions, and player counts.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### Bootstrap and Build
- **CRITICAL**: Always run commands from the repository root: `/home/runner/work/A2-Discord-Bot/A2-Discord-Bot`
- Download dependencies: `go mod download` -- takes ~1 second
- Verify module integrity: `go mod verify` -- takes <1 second  
- Tidy modules: `go mod tidy` -- takes <1 second
- Build the application: `go build -o a2bot main.go`
  - **First build**: ~13 seconds (downloading and compiling)  
  - **Subsequent builds**: <1 second (using build cache)
  - **NEVER CANCEL**: Set timeout to 30+ seconds for first builds
- Lint and format code:
  - `go vet ./...` -- checks for common errors, takes <1 second
  - `gofmt -d .` -- checks formatting (should show no output), takes <1 second

### Running the Application
- **Prerequisites**: Set environment variables `TOKEN` (Discord bot token) and `A2_API_KEY` (A2 station API key)
- Run with Go: `go run main.go` 
- Run compiled binary: `./a2bot`
- **Expected behavior without env vars**: Shows "Could not find token --> Set TOKEN in .env file." and exits cleanly
- Use provided scripts: 
  - `./start.sh` (Linux/Mac) - modify to set your environment variables
  - `start.bat` (Windows) - modify to set your environment variables

### Testing
- **No unit tests exist** in this repository
- **MANUAL VALIDATION REQUIRED**: Always test functionality changes using the validation script pattern
- Test similarity function: Input variations like "howmanyspacemonkey" match with 75% threshold
- Test table generation: Creates properly formatted ASCII tables with station data
- **Validation script example**:
```go
// Create /tmp/test_functionality.go to test components
package main
import ("fmt"; "a2-recreate/src")
func main() {
    // Test similarity
    fmt.Println(src.IsSimilar("howmanyspacemonkey", "howmanyspacemonke", 0.75)) // true
    // Test table generation with mock data
    mockFleets := []src.Fleet{...} // Add mock fleet data
    table := src.GenerateStationTable(mockFleets)
    fmt.Println(table)
}
```

### Docker Build
- **DOES NOT WORK**: Docker build fails due to network restrictions accessing Alpine package repositories
- Error occurs at `RUN apk add --no-cache git ca-certificates` step after ~60 seconds
- Do not attempt Docker builds - build times out after 2+ minutes with permission denied errors
- Use native Go build instead

## Validation
- **ALWAYS** run `go vet ./...` before committing - it must pass with no output (takes ~0.1 seconds)
- **ALWAYS** run `go build` to ensure code compiles successfully (takes ~0.4 seconds for clean builds)
- **MANUAL VALIDATION REQUIRED**: Test functionality using validation script pattern after making changes
- **CRITICAL VALIDATION SCENARIOS**:
  1. **Similarity Function**: Test `src.IsSimilar("variation", "howmanyspacemonke", 0.75)` returns expected boolean
  2. **Table Generation**: Test `src.GenerateStationTable()` with mock fleet data produces properly formatted ASCII table
  3. **API Format**: Ensure any changes to `src.GetAllFleets()` maintain expected JSON structure
  4. **Bot Startup**: Test `./a2bot` shows proper error message when environment variables missing
- The bot connects to Discord API and A2 station API, but cannot be fully tested without valid API credentials

## Common Tasks and Expected Outputs

### Repository Structure
```
.
├── README.md                 # Basic setup instructions
├── Dockerfile               # Docker build (fails due to network restrictions)
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums  
├── main.go                  # Main Discord bot logic
├── start.sh                 # Linux/Mac startup script
├── start.bat                # Windows startup script
├── assets/                  # Contains example bot response image
└── src/
    ├── a2fetchfleets.go     # A2 API client for fetching fleet data
    ├── simularitycheck.go   # Fuzzy string matching (Levenshtein distance)
    └── textartgenerator.go  # ASCII table generation
```

### Key Dependencies (from go.mod)
- `github.com/bwmarrin/discordgo v0.29.0` - Discord API client
- `github.com/joho/godotenv v1.5.1` - Environment variable loading
- Go 1.25.0 is required

### Functionality Overview
- **Trigger**: Bot responds to messages similar to "howmanyspacemonke" (75% similarity threshold)
- **Response**: Formatted ASCII table showing:
  - Station Name
  - Version  
  - Player Count
- **API**: Fetches data from `https://a2-station-api-prod-708695367983.us-central1.run.app/v2/fleets`
- **Similarity**: Uses Levenshtein distance algorithm for fuzzy command matching
- **Table**: Limits to 16 stations maximum, auto-sizes columns

### Environment Variables Required
- `TOKEN`: Discord bot token for authentication
- `A2_API_KEY`: API key for A2 station data service

### Common Issues and Solutions
- **Build errors**: Run `go mod tidy` then `go build` again
- **Vet warnings**: Format string issues in error messages - use `%w` for error wrapping, `%v` for value printing
- **Missing dependencies**: Run `go mod download` to fetch required modules
- **Docker build fails**: Use native Go build instead, Docker is not supported due to network restrictions

### Build and Test Timing Reference
- `go mod download`: ~0.01 seconds (when cached)
- `go mod verify`: ~0.08 seconds  
- `go mod tidy`: ~0.06 seconds
- `go vet ./...`: ~0.13 seconds
- `go build` (clean): ~0.35 seconds - NEVER CANCEL, set timeout to 10+ seconds
- `go build` (first time with downloads): ~13 seconds - NEVER CANCEL, set timeout to 30+ seconds  
- `go build` (cached): <0.4 seconds
- `go test ./...`: ~0.13 seconds (no tests exist, reports "no test files")
- Application startup test: <1 second (shows env var message and exits)
- Docker build: FAILS after 2+ minutes - DO NOT ATTEMPT

## Working with the Code
- **Entry point**: `main.go` contains Discord event handlers and bot initialization
- **Key functions**:
  - `src.IsSimilar()`: Fuzzy string matching for command detection
  - `src.GetAllFleets()`: Fetches fleet data from A2 API
  - `src.GenerateStationTable()`: Creates ASCII table from fleet data
- **Message handling**: Bot processes all messages and responds to similarity matches
- **API calls**: Authenticated with `A2_API_KEY` header to A2 station API

### Common Development Tasks
- **Adding new commands**: Modify the message handler in `main.go`, add similarity checks
- **Changing API endpoints**: Update URL in `src/a2fetchfleets.go`
- **Modifying table format**: Edit `src/textartgenerator.go`
- **Adjusting similarity threshold**: Change the 0.75 value in `main.go`
- **Adding new fields**: Update the `Station` struct in `src/a2fetchfleets.go` and table generation logic

### Error Handling Patterns
- Use `fmt.Errorf("message: %w", err)` for error wrapping
- Use `fmt.Printf("message %v", value)` for formatted output with values
- Always check `if err != nil` after API calls and operations
- API errors return proper error types that can be checked