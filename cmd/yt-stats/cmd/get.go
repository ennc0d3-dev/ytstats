package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

var (
	outputFormat string
	fields       []string
)

var getCmd = &cobra.Command{
	Use:   "get [video-id]",
	Short: "Get statistics for a YouTube video",
	Long: `Fetch and display statistics for a specific YouTube video by ID.

Examples:
  yt-stats get dQw4w9WgXcQ
  yt-stats get VIDEO_ID --format json
  yt-stats get VIDEO_ID --fields views,likes`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		videoID := args[0]
		getStats(videoID)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	
	getCmd.Flags().StringVarP(&outputFormat, "format", "f", "table", "Output format (table, json, yaml)")
	getCmd.Flags().StringSliceVar(&fields, "fields", []string{"views", "likes", "comments"}, "Fields to display")
}

func getStats(videoID string) {
	apiKey := viper.GetString("apiKey")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: YouTube API key is required")
		fmt.Fprintln(os.Stderr, "Set YTSTATS_API_KEY environment variable or use --api-key flag")
		os.Exit(1)
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating YouTube service: %v\n", err)
		os.Exit(1)
	}

	call := service.Videos.List([]string{"statistics", "snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching video stats: %v\n", err)
		os.Exit(1)
	}

	if len(response.Items) == 0 {
		fmt.Fprintf(os.Stderr, "No video found with ID: %s\n", videoID)
		os.Exit(1)
	}

	video := response.Items[0]
	stats := video.Statistics

	switch outputFormat {
	case "json":
		printJSON(video)
	case "yaml":
		printYAML(video)
	default:
		printTable(video, stats)
	}
}

func printTable(video *youtube.Video, stats *youtube.VideoStatistics) {
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Printf("║ %-57s ║\n", "YouTube Video Statistics")
	fmt.Println("╠═══════════════════════════════════════════════════════════╣")
	fmt.Printf("║ Title: %-50s ║\n", truncate(video.Snippet.Title, 50))
	fmt.Println("╠═══════════════════════════════════════════════════════════╣")
	
	for _, field := range fields {
		switch field {
		case "views", "viewCount":
			fmt.Printf("║ 👁️  Views:     %-43s ║\n", formatNumber(stats.ViewCount))
		case "likes", "likeCount":
			fmt.Printf("║ 👍 Likes:     %-43s ║\n", formatNumber(stats.LikeCount))
		case "comments", "commentCount":
			fmt.Printf("║ 💬 Comments:  %-43s ║\n", formatNumber(stats.CommentCount))
		case "favorites", "favoriteCount":
			fmt.Printf("║ ⭐ Favorites: %-43s ║\n", formatNumber(stats.FavoriteCount))
		}
	}
	
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
}

func printJSON(video *youtube.Video) {
	fmt.Printf(`{
  "id": "%s",
  "title": "%s",
  "viewCount": "%d",
  "likeCount": "%d",
  "commentCount": "%d",
  "favoriteCount": "%d"
}
`, video.Id, video.Snippet.Title, video.Statistics.ViewCount, video.Statistics.LikeCount, video.Statistics.CommentCount, video.Statistics.FavoriteCount)
}

func printYAML(video *youtube.Video) {
	fmt.Printf(`id: %s
title: %s
viewCount: %d
likeCount: %d
commentCount: %d
favoriteCount: %d
`, video.Id, video.Snippet.Title, video.Statistics.ViewCount, video.Statistics.LikeCount, video.Statistics.CommentCount, video.Statistics.FavoriteCount)
}

func formatNumber(n uint64) string {
	s := fmt.Sprintf("%d", n)
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}
