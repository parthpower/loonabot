// Package models
// JSON models for __a=1 responses
// special thanks to https://github.com/mholt/json-to-go/ for amazing tool
package models

type Insta struct {
	Items               []Items `json:"items"`
	NumResults          int     `json:"num_results"`
	MoreAvailable       bool    `json:"more_available"`
	AutoLoadMoreEnabled bool    `json:"auto_load_more_enabled"`
}
type Candidates struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	BloksAppURL               interface{} `json:"bloks_app_url"`
}
type CommentInformTreatment struct {
	ShouldHaveInformTreatment bool   `json:"should_have_inform_treatment"`
	Text                      string `json:"text"`
}
type CarouselMedia struct {
	ID                     string                 `json:"id"`
	MediaType              int                    `json:"media_type"`
	ImageVersions2         ImageVersions2         `json:"image_versions2"`
	OriginalWidth          int                    `json:"original_width"`
	OriginalHeight         int                    `json:"original_height"`
	VideoVersions          []VideoVersions        `json:"video_versions,omitempty"`
	VideoDuration          float64                `json:"video_duration,omitempty"`
	IsDashEligible         int                    `json:"is_dash_eligible,omitempty"`
	VideoDashManifest      string                 `json:"video_dash_manifest,omitempty"`
	VideoCodec             string                 `json:"video_codec,omitempty"`
	NumberOfQualities      int                    `json:"number_of_qualities,omitempty"`
	Pk                     int64                  `json:"pk"`
	CarouselParentID       string                 `json:"carousel_parent_id"`
	CanSeeInsightsAsBrand  bool                   `json:"can_see_insights_as_brand"`
	CommercialityStatus    string                 `json:"commerciality_status"`
	SharingFrictionInfo    SharingFrictionInfo    `json:"sharing_friction_info"`
	CommentInformTreatment CommentInformTreatment `json:"comment_inform_treatment"`
}
type FriendshipStatus struct {
	Following       bool `json:"following"`
	OutgoingRequest bool `json:"outgoing_request"`
	IsBestie        bool `json:"is_bestie"`
	IsRestricted    bool `json:"is_restricted"`
	IsFeedFavorite  bool `json:"is_feed_favorite"`
}
type User struct {
	Pk                         int64            `json:"pk"`
	Username                   string           `json:"username"`
	FullName                   string           `json:"full_name"`
	IsPrivate                  bool             `json:"is_private"`
	ProfilePicURL              string           `json:"profile_pic_url"`
	ProfilePicID               string           `json:"profile_pic_id"`
	FriendshipStatus           FriendshipStatus `json:"friendship_status"`
	IsVerified                 bool             `json:"is_verified"`
	FollowFrictionType         int              `json:"follow_friction_type"`
	HasAnonymousProfilePicture bool             `json:"has_anonymous_profile_picture"`
	IsUnpublished              bool             `json:"is_unpublished"`
	IsFavorite                 bool             `json:"is_favorite"`
	LatestReelMedia            int              `json:"latest_reel_media"`
	HasHighlightReels          bool             `json:"has_highlight_reels"`
	LiveBroadcastID            interface{}      `json:"live_broadcast_id"`
	LiveBroadcastVisibility    interface{}      `json:"live_broadcast_visibility"`
}
type MediaCroppingInfo struct {
	FeedPreviewCrop interface{} `json:"feed_preview_crop"`
	SquareCrop      interface{} `json:"square_crop"`
}
type Thumbnails struct {
	VideoLength                float64  `json:"video_length"`
	ThumbnailWidth             int      `json:"thumbnail_width"`
	ThumbnailHeight            int      `json:"thumbnail_height"`
	ThumbnailDuration          float64  `json:"thumbnail_duration"`
	SpriteUrls                 []string `json:"sprite_urls"`
	ThumbnailsPerRow           int      `json:"thumbnails_per_row"`
	TotalThumbnailNumPerSprite int      `json:"total_thumbnail_num_per_sprite"`
	MaxThumbnailsPerSprite     int      `json:"max_thumbnails_per_sprite"`
	SpriteWidth                int      `json:"sprite_width"`
	SpriteHeight               int      `json:"sprite_height"`
	RenderedWidth              int      `json:"rendered_width"`
	FileSizeKb                 int      `json:"file_size_kb"`
}

type IgtvFirstFrame struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
type FirstFrame struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
type AdditionalCandidates struct {
	IgtvFirstFrame IgtvFirstFrame `json:"igtv_first_frame"`
	FirstFrame     FirstFrame     `json:"first_frame"`
}
type Default struct {
	VideoLength                float64  `json:"video_length"`
	ThumbnailWidth             int      `json:"thumbnail_width"`
	ThumbnailHeight            int      `json:"thumbnail_height"`
	ThumbnailDuration          float64  `json:"thumbnail_duration"`
	SpriteUrls                 []string `json:"sprite_urls"`
	ThumbnailsPerRow           int      `json:"thumbnails_per_row"`
	TotalThumbnailNumPerSprite int      `json:"total_thumbnail_num_per_sprite"`
	MaxThumbnailsPerSprite     int      `json:"max_thumbnails_per_sprite"`
	SpriteWidth                int      `json:"sprite_width"`
	SpriteHeight               int      `json:"sprite_height"`
	RenderedWidth              int      `json:"rendered_width"`
	FileSizeKb                 int      `json:"file_size_kb"`
}
type AnimatedThumbnailSpritesheetInfoCandidates struct {
	Default Default `json:"default"`
}
type ImageVersions2 struct {
	Candidates                                 []Candidates                               `json:"candidates"`
	AdditionalCandidates                       AdditionalCandidates                       `json:"additional_candidates"`
	AnimatedThumbnailSpritesheetInfoCandidates AnimatedThumbnailSpritesheetInfoCandidates `json:"animated_thumbnail_spritesheet_info_candidates"`
}

type VideoVersions struct {
	Type   int    `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
	ID     string `json:"id"`
}
type Caption struct {
	Pk                 int64  `json:"pk"`
	UserID             int64  `json:"user_id"`
	Text               string `json:"text"`
	Type               int    `json:"type"`
	CreatedAt          int    `json:"created_at"`
	CreatedAtUtc       int    `json:"created_at_utc"`
	ContentType        string `json:"content_type"`
	Status             string `json:"status"`
	BitFlags           int    `json:"bit_flags"`
	DidReportAsSpam    bool   `json:"did_report_as_spam"`
	ShareEnabled       bool   `json:"share_enabled"`
	User               User   `json:"user"`
	IsCovered          bool   `json:"is_covered"`
	MediaID            int64  `json:"media_id"`
	PrivateReplyStatus int    `json:"private_reply_status"`
}

type IgArtist struct {
	Pk                 int64  `json:"pk"`
	Username           string `json:"username"`
	FullName           string `json:"full_name"`
	IsPrivate          bool   `json:"is_private"`
	ProfilePicURL      string `json:"profile_pic_url"`
	ProfilePicID       string `json:"profile_pic_id"`
	IsVerified         bool   `json:"is_verified"`
	FollowFrictionType int    `json:"follow_friction_type"`
}
type MusicMetadata struct {
	MusicCanonicalID  string      `json:"music_canonical_id"`
	AudioType         interface{} `json:"audio_type"`
	MusicInfo         interface{} `json:"music_info"`
	OriginalSoundInfo interface{} `json:"original_sound_info"`
}
type ConsumptionInfo struct {
	IsBookmarked          bool   `json:"is_bookmarked"`
	ShouldMuteAudioReason string `json:"should_mute_audio_reason"`
	IsTrendingInClips     bool   `json:"is_trending_in_clips"`
}
type OriginalSoundInfo struct {
	AudioAssetID                   int64           `json:"audio_asset_id"`
	ProgressiveDownloadURL         string          `json:"progressive_download_url"`
	DashManifest                   string          `json:"dash_manifest"`
	IgArtist                       IgArtist        `json:"ig_artist"`
	ShouldMuteAudio                bool            `json:"should_mute_audio"`
	OriginalMediaID                int64           `json:"original_media_id"`
	HideRemixing                   bool            `json:"hide_remixing"`
	DurationInMs                   int             `json:"duration_in_ms"`
	TimeCreated                    int             `json:"time_created"`
	OriginalAudioTitle             []string        `json:"original_audio_title"`
	ConsumptionInfo                ConsumptionInfo `json:"consumption_info"`
	AllowCreatorToRename           bool            `json:"allow_creator_to_rename"`
	CanRemixBeSharedToFb           bool            `json:"can_remix_be_shared_to_fb"`
	FormattedClipsMediaCount       interface{}     `json:"formatted_clips_media_count"`
	AudioParts                     []interface{}   `json:"audio_parts"`
	IsExplicit                     bool            `json:"is_explicit"`
	OriginalAudioSubtype           string          `json:"original_audio_subtype"`
	IsAudioAutomaticallyAttributed bool            `json:"is_audio_automatically_attributed"`
}
type MashupInfo struct {
	MashupsAllowed                      bool        `json:"mashups_allowed"`
	CanToggleMashupsAllowed             bool        `json:"can_toggle_mashups_allowed"`
	HasBeenMashedUp                     bool        `json:"has_been_mashed_up"`
	FormattedMashupsCount               interface{} `json:"formatted_mashups_count"`
	OriginalMedia                       interface{} `json:"original_media"`
	NonPrivacyFilteredMashupsMediaCount interface{} `json:"non_privacy_filtered_mashups_media_count"`
}
type BrandedContentTagInfo struct {
	CanAddTag bool `json:"can_add_tag"`
}
type AudioReattributionInfo struct {
	ShouldAllowRestore bool `json:"should_allow_restore"`
}
type AdditionalAudioInfo struct {
	AdditionalAudioUsername interface{}            `json:"additional_audio_username"`
	AudioReattributionInfo  AudioReattributionInfo `json:"audio_reattribution_info"`
}
type AudioRankingInfo struct {
	BestAudioClusterID string `json:"best_audio_cluster_id"`
}
type ClipsMetadata struct {
	MusicInfo                 interface{}           `json:"music_info"`
	OriginalSoundInfo         OriginalSoundInfo     `json:"original_sound_info"`
	AudioType                 string                `json:"audio_type"`
	MusicCanonicalID          string                `json:"music_canonical_id"`
	FeaturedLabel             interface{}           `json:"featured_label"`
	MashupInfo                MashupInfo            `json:"mashup_info"`
	NuxInfo                   interface{}           `json:"nux_info"`
	ViewerInteractionSettings interface{}           `json:"viewer_interaction_settings"`
	BrandedContentTagInfo     BrandedContentTagInfo `json:"branded_content_tag_info"`
	ShoppingInfo              interface{}           `json:"shopping_info"`
	AdditionalAudioInfo       AdditionalAudioInfo   `json:"additional_audio_info"`
	IsSharedToFb              bool                  `json:"is_shared_to_fb"`
	BreakingContentInfo       interface{}           `json:"breaking_content_info"`
	ChallengeInfo             interface{}           `json:"challenge_info"`
	ReelsOnTheRiseInfo        interface{}           `json:"reels_on_the_rise_info"`
	BreakingCreatorInfo       interface{}           `json:"breaking_creator_info"`
	AssetRecommendationInfo   interface{}           `json:"asset_recommendation_info"`
	ContextualHighlightInfo   interface{}           `json:"contextual_highlight_info"`
	ClipsCreationEntryPoint   string                `json:"clips_creation_entry_point"`
	AudioRankingInfo          AudioRankingInfo      `json:"audio_ranking_info"`
}
type SquareCrop struct {
	CropBottom float64 `json:"crop_bottom"`
	CropLeft   int     `json:"crop_left"`
	CropRight  int     `json:"crop_right"`
	CropTop    float64 `json:"crop_top"`
}

type Items struct {
	TakenAt                             int                    `json:"taken_at"`
	Pk                                  int64                  `json:"pk"`
	ID                                  string                 `json:"id"`
	DeviceTimestamp                     int64                  `json:"device_timestamp"`
	MediaType                           int                    `json:"media_type"`
	Code                                string                 `json:"code"`
	ClientCacheKey                      string                 `json:"client_cache_key"`
	FilterType                          int                    `json:"filter_type"`
	CarouselMediaCount                  int                    `json:"carousel_media_count"`
	CarouselMedia                       []CarouselMedia        `json:"carousel_media"`
	IsUnifiedVideo                      bool                   `json:"is_unified_video"`
	User                                User                   `json:"user"`
	CanViewerReshare                    bool                   `json:"can_viewer_reshare"`
	CaptionIsEdited                     bool                   `json:"caption_is_edited"`
	LikeAndViewCountsDisabled           bool                   `json:"like_and_view_counts_disabled"`
	FeaturedProductsCta                 interface{}            `json:"featured_products_cta"`
	CommercialityStatus                 string                 `json:"commerciality_status"`
	IsPaidPartnership                   bool                   `json:"is_paid_partnership"`
	IsVisualReplyCommenterNoticeEnabled bool                   `json:"is_visual_reply_commenter_notice_enabled"`
	OriginalMediaHasVisualReplyMedia    bool                   `json:"original_media_has_visual_reply_media"`
	CommentsDisabled                    bool                   `json:"comments_disabled"`
	CommentLikesEnabled                 bool                   `json:"comment_likes_enabled"`
	CommentThreadingEnabled             bool                   `json:"comment_threading_enabled"`
	HasMoreComments                     bool                   `json:"has_more_comments"`
	MaxNumVisiblePreviewComments        int                    `json:"max_num_visible_preview_comments"`
	PreviewComments                     []interface{}          `json:"preview_comments"`
	Comments                            []interface{}          `json:"comments"`
	CanViewMorePreviewComments          bool                   `json:"can_view_more_preview_comments"`
	CommentCount                        int                    `json:"comment_count"`
	HideViewAllCommentEntrypoint        bool                   `json:"hide_view_all_comment_entrypoint"`
	InlineComposerDisplayCondition      string                 `json:"inline_composer_display_condition"`
	Title                               string                 `json:"title"`
	NearlyCompleteCopyrightMatch        bool                   `json:"nearly_complete_copyright_match"`
	Thumbnails                          Thumbnails             `json:"thumbnails"`
	IgtvExistsInViewerSeries            bool                   `json:"igtv_exists_in_viewer_series"`
	IsPostLive                          bool                   `json:"is_post_live"`
	ImageVersions2                      ImageVersions2         `json:"image_versions2"`
	OriginalWidth                       int                    `json:"original_width"`
	OriginalHeight                      int                    `json:"original_height"`
	LikeCount                           int                    `json:"like_count"`
	HasLiked                            bool                   `json:"has_liked"`
	TopLikers                           []interface{}          `json:"top_likers"`
	FacepileTopLikers                   []interface{}          `json:"facepile_top_likers"`
	PhotoOfYou                          bool                   `json:"photo_of_you"`
	CanSeeInsightsAsBrand               bool                   `json:"can_see_insights_as_brand"`
	MashupInfo                          MashupInfo             `json:"mashup_info"`
	IsDashEligible                      int                    `json:"is_dash_eligible"`
	VideoDashManifest                   string                 `json:"video_dash_manifest"`
	VideoCodec                          string                 `json:"video_codec"`
	NumberOfQualities                   int                    `json:"number_of_qualities"`
	VideoVersions                       []VideoVersions        `json:"video_versions"`
	HasAudio                            bool                   `json:"has_audio"`
	VideoDuration                       float64                `json:"video_duration"`
	ViewCount                           int                    `json:"view_count"`
	PlayCount                           int                    `json:"play_count"`
	Caption                             Caption                `json:"caption"`
	CanViewerSave                       bool                   `json:"can_viewer_save"`
	OrganicTrackingToken                string                 `json:"organic_tracking_token"`
	SharingFrictionInfo                 SharingFrictionInfo    `json:"sharing_friction_info"`
	CommentInformTreatment              CommentInformTreatment `json:"comment_inform_treatment"`
	ProductType                         string                 `json:"product_type"`
	IsInProfileGrid                     bool                   `json:"is_in_profile_grid"`
	ProfileGridControlEnabled           bool                   `json:"profile_grid_control_enabled"`
	DeletedReason                       int                    `json:"deleted_reason"`
	IntegrityReviewDecision             string                 `json:"integrity_review_decision"`
	MusicMetadata                       interface{}            `json:"music_metadata"`
	ClipsMetadata                       ClipsMetadata          `json:"clips_metadata"`
	MediaCroppingInfo                   MediaCroppingInfo      `json:"media_cropping_info"`
}

const (
	MediaTypeImage    = 1
	MediaTypeVideo    = 2
	MediaTypeCarousel = 8
)
