package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// BenchmarkResult stores benchmark results for a RAID level
type BenchmarkResult struct {
	RaidType       string
	WriteTime      float64
	WriteSpeed     float64
	ReadTime       float64
	ReadSpeed      float64
	EffectiveCap   int
	OverheadPct    float64
}

// PrintBarChart creates a simple ASCII bar chart
func PrintBarChart(title string, data map[string]float64, maxWidth int) {
	fmt.Printf("\n%s:\n", title)
	
	// Find the maximum value
	maxVal := 0.0
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}
	
	// Print the chart
	for name, val := range data {
		barWidth := int((val / maxVal) * float64(maxWidth))
		bar := strings.Repeat("█", barWidth)
		fmt.Printf("%-6s [%s] %.2f\n", name, bar, val)
	}
}

// SaveResultsToFile saves benchmark results to a file for later analysis
func SaveResultsToFile(results []BenchmarkResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Write header
	_, err = file.WriteString("RaidType,WriteTime,WriteSpeed,ReadTime,ReadSpeed,EffectiveCap,OverheadPct\n")
	if err != nil {
		return err
	}
	
	// Write data
	for _, result := range results {
		line := fmt.Sprintf("%s,%.2f,%.2f,%.2f,%.2f,%d,%.2f\n",
			result.RaidType, 
			result.WriteTime, 
			result.WriteSpeed, 
			result.ReadTime, 
			result.ReadSpeed, 
			result.EffectiveCap, 
			result.OverheadPct)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}
	
	return nil
}

// VisualizeResults creates ASCII-based visualizations of the benchmark results
func VisualizeResults(results []BenchmarkResult) {
	// Create maps for different metrics
	writeSpeed := make(map[string]float64)
	readSpeed := make(map[string]float64)
	effectiveCapacity := make(map[string]float64)
	overhead := make(map[string]float64)
	
	for _, result := range results {
		writeSpeed[result.RaidType] = result.WriteSpeed
		readSpeed[result.RaidType] = result.ReadSpeed
		effectiveCapacity[result.RaidType] = float64(result.EffectiveCap)
		overhead[result.RaidType] = result.OverheadPct
	}
	
	// Print bar charts
	const maxWidth = 50
	PrintBarChart("Write Speed (MB/s)", writeSpeed, maxWidth)
	PrintBarChart("Read Speed (MB/s)", readSpeed, maxWidth)
	PrintBarChart("Effective Capacity (MB)", effectiveCapacity, maxWidth)
	PrintBarChart("Storage Overhead (%)", overhead, maxWidth)
	
	// Save results to file
	err := SaveResultsToFile(results, "raid_benchmark_results.csv")
	if err != nil {
		log.Printf("Error saving results to file: %v", err)
	} else {
		fmt.Println("\nResults saved to raid_benchmark_results.csv")
	}
}

// ParseMainOutput parses the output from the main program to extract benchmark results
func ParseMainOutput(output string) []BenchmarkResult {
	var results []BenchmarkResult
	
	lines := strings.Split(output, "\n")
	dataStarted := false
	
	for _, line := range lines {
		// Skip lines until we find the header row
		if strings.Contains(line, "RAID") && strings.Contains(line, "Write Time") {
			dataStarted = true
			continue
		}
		
		// Process data rows
		if dataStarted && len(line) > 0 && strings.HasPrefix(line, "RAID") {
			fields := strings.Fields(line)
			if len(fields) >= 7 {
				// Parse fields
				raidType := fields[0]
				writeTime, _ := strconv.ParseFloat(strings.Replace(fields[1], "seconds", "", -1), 64)
				writeSpeed, _ := strconv.ParseFloat(fields[2], 64)
				readTime, _ := strconv.ParseFloat(strings.Replace(fields[3], "seconds", "", -1), 64)
				readSpeed, _ := strconv.ParseFloat(fields[4], 64)
				effectiveCap, _ := strconv.Atoi(fields[5])
				overhead, _ := strconv.ParseFloat(fields[6], 64)
				
				results = append(results, BenchmarkResult{
					RaidType:     raidType,
					WriteTime:    writeTime,
					WriteSpeed:   writeSpeed,
					ReadTime:     readTime,
					ReadSpeed:    readSpeed,
					EffectiveCap: effectiveCap,
					OverheadPct:  overhead,
				})
			}
		}
		
		// Stop processing when we reach the analysis section
		if dataStarted && strings.Contains(line, "Analysis:") {
			break
		}
	}
	
	return results
}

func main() {
	// Example usage:
	// 1. Run the main RAID benchmark
	// 2. Capture its output
	// 3. Parse the output
	// 4. Visualize the results
	
	// For demonstration, we'll use some sample results
	results := []BenchmarkResult{
		{RaidType: "RAID0", WriteTime: 1.23, WriteSpeed: 81.30, ReadTime: 0.42, ReadSpeed: 238.10, EffectiveCap: 500, OverheadPct: 0.00},
		{RaidType: "RAID1", WriteTime: 5.78, WriteSpeed: 17.30, ReadTime: 0.40, ReadSpeed: 250.00, EffectiveCap: 100, OverheadPct: 80.00},
		{RaidType: "RAID4", WriteTime: 3.24, WriteSpeed: 30.86, ReadTime: 0.48, ReadSpeed: 208.33, EffectiveCap: 400, OverheadPct: 20.00},
		{RaidType: "RAID5", WriteTime: 2.98, WriteSpeed: 33.56, ReadTime: 0.45, ReadSpeed: 222.22, EffectiveCap: 400, OverheadPct: 20.00},
	}
	
	// Visualize the results
	VisualizeResults(results)
	
	fmt.Println("\nNote: This visualization script can be used after running the main RAID benchmark.")
	fmt.Println("To use it with actual benchmark results, you would need to capture the output")
	fmt.Println("from the main program and pass it to ParseMainOutput() function.")
}