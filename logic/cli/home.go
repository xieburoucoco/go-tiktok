package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kaptinlin/jsonrepair"
	"go-tiktok/consts"
	"regexp"
)

func (s *STikTokCli) UnmarshalHomeResponse(ctx context.Context, body []byte) (HomeJson, error) {
	res := HomeJson{}
	matchArr := regexp.MustCompile(consts.TIKTOK_HOME_EXTRACT_PATTERN).FindAllStringSubmatch(string(body), -1)
	if len(matchArr) == 0 {
		return res, fmt.Errorf("No matching string was found, parsing failed")
	}
	matchStr := matchArr[0][1]
	repair, err := jsonrepair.JSONRepair(matchStr)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(repair), &res)
	return res, err
}

type HomeJson struct {
	Default struct {
		WebAppContext struct {
			Language string `json:"language"`
			Region   string `json:"region"`
			AppID    int    `json:"appId"`
			AppType  string `json:"appType"`
			Wid      string `json:"wid"`
			//WebIDCreatedTime string `json:"webIdCreatedTime"`
			OdinID        string `json:"odinId"`
			Nonce         string `json:"nonce"`
			BotType       string `json:"botType"`
			RequestID     string `json:"requestId"`
			ClusterRegion string `json:"clusterRegion"`
			AbTestVersion struct {
				VersionName string `json:"versionName"`
				Parameters  struct {
					MobileConsumptionLimitLogin struct {
						Vid string `json:"vid"`
					} `json:"mobile_consumption_limit_login"`
					MobileConsumptionLimitLoggedIn struct {
						Vid string `json:"vid"`
					} `json:"mobile_consumption_limit_logged_in"`
					UseInboxNoticeCountAPI struct {
						Vid string `json:"vid"`
					} `json:"use_inbox_notice_count_api"`
					OneColumnPlayerSize struct {
						Vid string `json:"vid"`
					} `json:"one_column_player_size"`
					VideoServerpush struct {
						Vid string `json:"vid"`
					} `json:"video_serverpush"`
					QrSsoPopup struct {
						Vid string `json:"vid"`
					} `json:"qr_sso_popup"`
					ShouldHighlightHashtag struct {
						Vid string `json:"vid"`
					} `json:"should_highlight_hashtag"`
					BrowserLoginRedirect struct {
						Vid string `json:"vid"`
					} `json:"browser_login_redirect"`
					WebappSwitchAccount struct {
						Vid string `json:"vid"`
					} `json:"webapp_switch_account"`
					MobileVodkit struct {
						Vid string `json:"vid"`
					} `json:"mobile_vodkit"`
					XgVolumeTest struct {
						Vid string `json:"vid"`
					} `json:"xg_volume_test"`
					VideoFeedRedesign struct {
						Vid string `json:"vid"`
					} `json:"video_feed_redesign"`
					VideoBitrateAdapt struct {
						Vid string `json:"vid"`
					} `json:"video_bitrate_adapt"`
					MobileConsumptionLimitNonLoggedIn struct {
						Vid string `json:"vid"`
					} `json:"mobile_consumption_limit_non_logged_in"`
					SearchVideo struct {
						Vid    string `json:"vid"`
						BotVid string `json:"botVid"`
					} `json:"search_video"`
					RemoveBottomBanner struct {
						Vid string `json:"vid"`
					} `json:"remove_bottom_banner"`
					MobilePredictiveData struct {
						Vid string `json:"vid"`
					} `json:"mobile_predictive_data"`
					PeriodicLoginPopupInterval struct {
						Vid string `json:"vid"`
					} `json:"periodic_login_popup_interval"`
					EnableMlModel struct {
						Vid string `json:"vid"`
					} `json:"enable_ml_model"`
					LoginModalUIRevamp struct {
						Vid string `json:"vid"`
					} `json:"login_modal_ui_revamp"`
					NonLoggedInComments struct {
						Vid string `json:"vid"`
					} `json:"non_logged_in_comments"`
					VolumeNormalize struct {
						Vid string `json:"vid"`
					} `json:"volume_normalize"`
					MobileSearchTest struct {
						Vid string `json:"vid"`
					} `json:"mobile_search_test"`
					ShareButtonPart1Test struct {
						Vid string `json:"vid"`
					} `json:"share_button_part1_test"`
					ConfirmLogout struct {
						Vid string `json:"vid"`
					} `json:"confirm_logout"`
					LoginOptionOrderByMetrics struct {
						Vid string `json:"vid"`
					} `json:"login_option_order_by_metrics"`
					LastLoginMethod struct {
						Vid string `json:"vid"`
					} `json:"last_login_method"`
					BrowserModeEncourageLogin struct {
						Vid string `json:"vid"`
					} `json:"browser_mode_encourage_login"`
					LoginModalImage struct {
						Vid string `json:"vid"`
					} `json:"login_modal_image"`
					CreatorCenterConnect struct {
						Vid string `json:"vid"`
					} `json:"creator_center_connect"`
					WebappLoginEmailPhone struct {
						Vid string `json:"vid"`
					} `json:"webapp_login_email_phone"`
					XgplayerPreloadConfig struct {
						Vid string `json:"vid"`
					} `json:"xgplayer_preload_config"`
					SignUpWebappRegionChange struct {
						Vid string `json:"vid"`
					} `json:"sign_up_webapp_region_change"`
					UseFollowV2 struct {
						Vid string `json:"vid"`
					} `json:"use_follow_v2"`
					ZTITest struct {
						ConsumerPathList []string `json:"consumer_path_list"`
						Vid              string   `json:"vid"`
					} `json:"ZTI_test"`
					AbTag struct {
						MergeRiskEvent int `json:"merge_risk_event"`
					} `json:"ab_tag"`
					AbTags struct {
						PrefetchVisaCryptogram bool `json:"prefetch_visa_cryptogram"`
					} `json:"ab_tags"`
					AddKapEntry struct {
						Vid string `json:"vid"`
					} `json:"add_kap_entry"`
					AddProfileLeftBar struct {
						Vid string `json:"vid"`
					} `json:"add_profile_left_bar"`
					AddTranscript struct {
						Vid string `json:"vid"`
					} `json:"add_transcript"`
					AnalyticsUpgradePhase3 struct {
						Vid string `json:"vid"`
					} `json:"analytics_upgrade_phase3"`
					AutoDarkMode struct {
						Vid string `json:"vid"`
					} `json:"auto_dark_mode"`
					AutoScroll struct {
						Vid string `json:"vid"`
					} `json:"auto_scroll"`
					AutoscrollReposition struct {
						Vid string `json:"vid"`
					} `json:"autoscroll_reposition"`
					BrowserModeCreatorTab3 struct {
						Vid string `json:"vid"`
					} `json:"browser_mode_creator_tab_3"`
					CcPerfPhase1 struct {
						Vid struct {
							IsCreatorCenterContextEnabled   bool `json:"isCreatorCenterContextEnabled"`
							IsInsightV2Enabled              bool `json:"isInsightV2Enabled"`
							IsOfflineI18NEnabled            bool `json:"isOfflineI18nEnabled"`
							IsParallelIframeEnabled         bool `json:"isParallelIframeEnabled"`
							IsPhase2Enabled                 bool `json:"isPhase2Enabled"`
							IsPrefetchIframeResourceEnabled bool `json:"isPrefetchIframeResourceEnabled"`
							IsServerSideTranslationEnabled  bool `json:"isServerSideTranslationEnabled"`
						} `json:"vid"`
					} `json:"cc_perf_phase1"`
					ChangeListLengthNew struct {
						Vid string `json:"vid"`
					} `json:"change_list_length_new"`
					CleanFixedBottom struct {
						Vid string `json:"vid"`
					} `json:"clean_fixed_bottom"`
					CommentRefactorTest struct {
						Vid string `json:"vid"`
					} `json:"comment_refactor_test"`
					CreatorCenterConnectGlobal struct {
						Vid string `json:"vid"`
					} `json:"creator_center_connect_global"`
					CreatorCenterGlobalCommentManagement struct {
						Vid string `json:"vid"`
					} `json:"creator_center_global_comment_management"`
					CreatorCenterTest struct {
						Vid string `json:"vid"`
					} `json:"creator_center_test"`
					DelayGuest struct {
						Vid string `json:"vid"`
					} `json:"delay_guest"`
					DelayGuestOthersHomepage struct {
						Vid string `json:"vid"`
					} `json:"delay_guest_others_homepage"`
					DesktopUIOpt struct {
						Vid string `json:"vid"`
					} `json:"desktop_ui_opt"`
					DesktopUIReply struct {
						Vid string `json:"vid"`
					} `json:"desktop_ui_reply"`
					DetailPageCommentsRedesign struct {
						Vid string `json:"vid"`
					} `json:"detail_page_comments_redesign"`
					DigitalWellbeingWeb struct {
						Vid string `json:"vid"`
					} `json:"digital_wellbeing_web"`
					EnableAboutThisAd struct {
						Vid string `json:"vid"`
					} `json:"enable_about_this_ad"`
					EnableAds struct {
						Vid string `json:"vid"`
					} `json:"enable_ads"`
					EnableAutoscrollMoremenu struct {
						Vid string `json:"vid"`
					} `json:"enable_autoscroll_moremenu"`
					EnableContinuePlay struct {
						Vid string `json:"vid"`
					} `json:"enable_continue_play"`
					EnableEcLcc struct {
						Vid string `json:"vid"`
					} `json:"enable_ec_lcc"`
					EnableFbSdk struct {
						Vid string `json:"vid"`
					} `json:"enable_fb_sdk"`
					EnableMiniPlayer struct {
						Vid string `json:"vid"`
					} `json:"enable_mini_player"`
					EnableMiniPlayerNewDesign struct {
						Vid string `json:"vid"`
					} `json:"enable_mini_player_new_design"`
					EnableMiniPlayerSwitchTabPopup struct {
						Vid string `json:"vid"`
					} `json:"enable_mini_player_switch_tab_popup"`
					EnableNotInterested struct {
						Vid string `json:"vid"`
					} `json:"enable_not_interested"`
					EnableOdinID struct {
						Vid string `json:"vid"`
					} `json:"enable_odin_id"`
					EnableProfilePinnedVideo struct {
						Vid string `json:"vid"`
					} `json:"enable_profile_pinned_video"`
					EnhanceVideoConsumptionTest struct {
						Vid string `json:"vid"`
					} `json:"enhance_video_consumption_test"`
					ExchangeRetentionPopup struct {
						Vid string `json:"vid"`
					} `json:"exchange_retention_popup"`
					ExpandItemTag struct {
						Vid string `json:"vid"`
					} `json:"expand_item_tag"`
					ExploreAllTab struct {
						Vid string `json:"vid"`
					} `json:"explore_all_tab"`
					ExploreShuntTest struct {
						Vid string `json:"vid"`
					} `json:"explore_shunt_test"`
					ExploreTest struct {
						Vid string `json:"vid"`
					} `json:"explore_test"`
					ExploreUIChange struct {
						Vid string `json:"vid"`
					} `json:"explore_ui_change"`
					FavoriteTest struct {
						Vid string `json:"vid"`
					} `json:"favorite_test"`
					FixTeaSession struct {
						Vid string `json:"vid"`
					} `json:"fix_tea_session"`
					FollowingDisplayLive struct {
						Vid string `json:"vid"`
					} `json:"following_display_live"`
					FriendsTab struct {
						Vid string `json:"vid"`
					} `json:"friends_tab"`
					FypBackupV2 struct {
						Enable string `json:"enable"`
					} `json:"fyp_backup_v2"`
					FypCommentsPanel struct {
						Vid string `json:"vid"`
					} `json:"fyp_comments_panel"`
					FypProgressBar struct {
						Vid string `json:"vid"`
					} `json:"fyp_progress_bar"`
					FypRedesign struct {
						Vid string `json:"vid"`
					} `json:"fyp_redesign"`
					FypTopGradient struct {
						Vid string `json:"vid"`
					} `json:"fyp_top_gradient"`
					FypUpdateGradient struct {
						Vid string `json:"vid"`
					} `json:"fyp_update_gradient"`
					GuestModeRedesign struct {
						Vid string `json:"vid"`
					} `json:"guest_mode_redesign"`
					IncreaseDetailPageCoverQuantityTest struct {
						Vid string `json:"vid"`
					} `json:"increase_detail_page_cover_quantity_test"`
					IslandsArchExplore struct {
						Vid string `json:"vid"`
					} `json:"islands_arch_explore"`
					IslandsArchPhase2 struct {
						Vid string `json:"vid"`
					} `json:"islands_arch_phase2"`
					IslandsArchVideoDetail struct {
						Vid string `json:"vid"`
					} `json:"islands_arch_video_detail"`
					KepNewUILogin struct {
						Vid string `json:"vid"`
					} `json:"kep_new_ui_login"`
					KepVideoSortCtrExp struct {
						Vid string `json:"vid"`
					} `json:"kep_video_sort_ctr_exp"`
					KepVideos struct {
						PredictParams struct {
							ValueTree struct {
								Base struct {
									RankIDWeight float64 `json:"rank_id_weight"`
								} `json:"base"`
							} `json:"value_tree"`
							Formulas struct {
								CoverAnsaOk  string `json:"cover_ansa_ok"`
								ReportOk     string `json:"report_ok"`
								MultimodalOk string `json:"multimodal_ok"`
								IsStyleOk    string `json:"is_style_ok"`
								CmBoost      string `json:"cm_boost"`
								RankIDScore  string `json:"rank_id_score"`
								AbSlot1      string `json:"ab_slot_1"`
								AbSlot2      string `json:"ab_slot_2"`
							} `json:"formulas"`
						} `json:"predict_params"`
					} `json:"kep_videos"`
					LiveAbrVersion struct {
						Vid string `json:"vid"`
					} `json:"live_abr_version"`
					LiveContextOptimize struct {
						Vid string `json:"vid"`
					} `json:"live_context_optimize"`
					LiveCsrInsertContext struct {
						Vid string `json:"vid"`
					} `json:"live_csr_insert_context"`
					LiveDetailNonLoggedInEntry struct {
						Seconds int    `json:"seconds"`
						Vid     string `json:"vid"`
					} `json:"live_detail_non_logged_in_entry"`
					LiveEndImprovedMetrics struct {
						Vid string `json:"vid"`
					} `json:"live_end_improved_metrics"`
					LiveEnterRoomChore struct {
						Vid string `json:"vid"`
					} `json:"live_enter_room_chore"`
					LiveEventAggregation struct {
						Vid string `json:"vid"`
					} `json:"live_event_aggregation"`
					LiveFeedStyle struct {
						Vid string `json:"vid"`
					} `json:"live_feed_style"`
					LiveGoliveEntrance struct {
						Vid string `json:"vid"`
					} `json:"live_golive_entrance"`
					LiveHeaderDelay struct {
						Vid string `json:"vid"`
					} `json:"live_header_delay"`
					LiveI18NReduce struct {
						Vid string `json:"vid"`
					} `json:"live_i18n_reduce"`
					LiveImStarlingCdn struct {
						Vid string `json:"vid"`
					} `json:"live_im_starling_cdn"`
					LiveLcpPerfOptimize struct {
						Vid string `json:"vid"`
					} `json:"live_lcp_perf_optimize"`
					LiveLike struct {
						Vid string `json:"vid"`
					} `json:"live_like"`
					LiveLoginReflowBtn struct {
						Vid string `json:"vid"`
					} `json:"live_login_reflow_btn"`
					LiveLowLatency struct {
						Vid string `json:"vid"`
					} `json:"live_low_latency"`
					LiveNewDiscover struct {
						Vid string `json:"vid"`
					} `json:"live_new_discover"`
					LivePlayerIcon struct {
						Vid string `json:"vid"`
					} `json:"live_player_icon"`
					LivePlayerMuteText struct {
						Vid string `json:"vid"`
					} `json:"live_player_mute_text"`
					LivePlayerSwitchButton struct {
						Vid string `json:"vid"`
					} `json:"live_player_switch_button"`
					LivePreviewWeb struct {
						Vid string `json:"vid"`
					} `json:"live_preview_web"`
					LiveProShow struct {
						Vid string `json:"vid"`
					} `json:"live_pro_show"`
					LiveRechargeByAmount struct {
						Vid string `json:"vid"`
					} `json:"live_recharge_by_amount"`
					LiveRechargeHomescreen struct {
						Vid string `json:"vid"`
					} `json:"live_recharge_homescreen"`
					LiveRoomAgeRestriction struct {
						Vid string `json:"vid"`
					} `json:"live_room_age_restriction"`
					LiveRoomCsr struct {
						Vid string `json:"vid"`
					} `json:"live_room_csr"`
					LiveRoomMatch struct {
						Vid string `json:"vid"`
					} `json:"live_room_match"`
					LiveRoomNonStreaming struct {
						Vid string `json:"vid"`
					} `json:"live_room_non_streaming"`
					LiveStageDefinition struct {
						Vid string `json:"vid"`
					} `json:"live_stage_definition"`
					LiveStudioDownloadRefactorPc struct {
						Vid string `json:"vid"`
					} `json:"live_studio_download_refactor_pc"`
					LiveTopViewers struct {
						Vid string `json:"vid"`
					} `json:"live_top_viewers"`
					LiveWalletPerformancePackup struct {
						Vid string `json:"vid"`
					} `json:"live_wallet_performance_packup"`
					LiveWebcodecsPlayer struct {
						Vid string `json:"vid"`
					} `json:"live_webcodecs_player"`
					NewItemTag struct {
						Vid string `json:"vid"`
					} `json:"new_item_tag"`
					NewServerCall struct {
						Vid string `json:"vid"`
					} `json:"new_server_call"`
					NewTransparencyCenter struct {
						Vid string `json:"vid"`
					} `json:"new_transparency_center"`
					OptimiseBrowserMode struct {
						Vid string `json:"vid"`
					} `json:"optimise_browser_mode"`
					PcInspiration struct {
						Vid string `json:"vid"`
					} `json:"pc_inspiration"`
					PcVideoPlaylistTest struct {
						Vid string `json:"vid"`
					} `json:"pc_video_playlist_test"`
					PhotoModeYml struct {
						Vid string `json:"vid"`
					} `json:"photo_mode_yml"`
					PhotoTest struct {
						Vid string `json:"vid"`
					} `json:"photo_test"`
					PlayerDegrade struct {
						Vid string `json:"vid"`
					} `json:"player_degrade"`
					PnsKeywordFiltering struct {
						Vid string `json:"vid"`
					} `json:"pns_keyword_filtering"`
					PnsPopupSdk struct {
						Vid string `json:"vid"`
					} `json:"pns_popup_sdk"`
					PreviewCover struct {
						Vid string `json:"vid"`
					} `json:"preview_cover"`
					ProfileFollowInfo struct {
						Vid string `json:"vid"`
					} `json:"profile_follow_info"`
					PromoteQrCode struct {
						Vid string `json:"vid"`
					} `json:"promote_qr_code"`
					ReduceUserItemList struct {
						Vid string `json:"vid"`
					} `json:"reduce_user_item_list"`
					RefactorAnchorTag struct {
						Vid string `json:"vid"`
					} `json:"refactor_anchor_tag"`
					RelatedVideosVideoDetailPage struct {
						Vid string `json:"vid"`
					} `json:"related_videos_video_detail_page"`
					RemoveDisclaimer struct {
						Vid string `json:"vid"`
					} `json:"remove_disclaimer"`
					RemoveTooltip struct {
						Vid string `json:"vid"`
					} `json:"remove_tooltip"`
					ReportItemTag struct {
						Vid string `json:"vid"`
					} `json:"report_item_tag"`
					ReverseExpandItemTag struct {
						Vid string `json:"vid"`
					} `json:"reverse_expand_item_tag"`
					SearchAddLive struct {
						Vid string `json:"vid"`
					} `json:"search_add_live"`
					SearchAddRelatedSearch struct {
						Vid string `json:"vid"`
					} `json:"search_add_related_search"`
					SearchBarStyleOpt struct {
						Vid string `json:"vid"`
					} `json:"search_bar_style_opt"`
					SearchKeepSugShow struct {
						Vid string `json:"vid"`
					} `json:"search_keep_sug_show"`
					SearchTopAuthorCard struct {
						Vid string `json:"vid"`
					} `json:"search_top_author_card"`
					SearchTransferHistory struct {
						Vid string `json:"vid"`
					} `json:"search_transfer_history"`
					SearchVideoLab struct {
						Vid string `json:"vid"`
					} `json:"search_video_lab"`
					SeoBreadcrumbDetail struct {
						Vid string `json:"vid"`
					} `json:"seo_breadcrumb_detail"`
					SeoDesktop struct {
						Vid string `json:"vid"`
					} `json:"seo_desktop"`
					ShouldRecomReduceIconRisk struct {
						Vid string `json:"vid"`
					} `json:"should_recom_reduce_icon_risk"`
					ShowAigcLabelWeb struct {
						Vid string `json:"vid"`
					} `json:"show_aigc_label_web"`
					SidenavTest struct {
						Vid string `json:"vid"`
					} `json:"sidenav_test"`
					StudioWebEhEntrance struct {
						Vid string `json:"vid"`
					} `json:"studio_web_eh_entrance"`
					StudioWebEhEntranceV2 struct {
						Vid string `json:"vid"`
					} `json:"studio_web_eh_entrance_v2"`
					ThreeColumnGif struct {
						Vid string `json:"vid"`
					} `json:"three_column_gif"`
					Tikcast struct {
						EnableEcLcc        bool `json:"enable_ec_lcc"`
						LiveRestrictedMode bool `json:"live_restricted_mode"`
					} `json:"tikcast"`
					Tiktok struct {
						LiveRestrictedMode         bool `json:"live_restricted_mode"`
						PrivateAccountPromptForU18 int  `json:"private_account_prompt_for_u18"`
						UseAlignedCopies           int  `json:"use_aligned_copies"`
					} `json:"tiktok"`
					TiktokWeb struct {
						CapcutEntryGroup                   int    `json:"capcut_entry_group"`
						CopyrightPrecheck                  string `json:"copyright_precheck"`
						EnableCloudDraft                   string `json:"enable_cloud_draft"`
						EnableLocalDraft                   string `json:"enable_local_draft"`
						EnableNewPlaylist                  string `json:"enable_new_playlist"`
						EnableNewPostAPI                   string `json:"enable_new_post_api"`
						EnableWeb60MinFlag                 bool   `json:"enable_web_60_min_flag"`
						MaxDividedNum                      string `json:"max_divided_num"`
						OptV1NewEntranceV3                 int    `json:"opt_v1_new_entrance_v3"`
						OptV1NewUploadUI                   int    `json:"opt_v1_new_upload_ui"`
						PostPollingVersion                 int    `json:"post_polling_version"`
						ShowAigcToggle                     int    `json:"show_aigc_toggle"`
						ShowCreatorAcademyPcEntrance       int    `json:"show_creator_academy_pc_entrance"`
						TtsProductAnchor                   int    `json:"tts_product_anchor"`
						UploadFileStepOptimization         int    `json:"upload_file_step_optimization"`
						UseVideoPreviewTranscodeMultiChunk bool   `json:"use_video_preview_transcode_multi_chunk"`
						VideoSplitCountLimit               int    `json:"video_split_count_limit"`
						WebCreationCoverTool               int    `json:"web_creation_cover_tool"`
						WebCreationLocalVideoPreview       int    `json:"web_creation_local_video_preview"`
						WebCreationSupportEdit             int    `json:"web_creation_support_edit"`
					} `json:"tiktok_web"`
					TranslationReduce struct {
						Vid string `json:"vid"`
					} `json:"translation_reduce"`
					TtehEffectAnchorV1 struct {
						Vid string `json:"vid"`
					} `json:"tteh_effect_anchor_v1"`
					TtliveBroadcastTopicVersionTwo struct {
						Vid string `json:"vid"`
					} `json:"ttlive_broadcast_topic_version_two"`
					UILayoutAlignment struct {
						Vid string `json:"vid"`
					} `json:"ui_layout_alignment"`
					UseAlignedCopies struct {
						Vid string `json:"vid"`
					} `json:"use_aligned_copies"`
					UseErrorBoundary struct {
						Vid string `json:"vid"`
					} `json:"use_error_boundary"`
					UseLeftNavigationRefactor struct {
						Vid string `json:"vid"`
					} `json:"use_left_navigation_refactor"`
					UseNavigationRefactor struct {
						Vid string `json:"vid"`
					} `json:"use_navigation_refactor"`
					VideoDetailAuthorCard struct {
						Vid string `json:"vid"`
					} `json:"video_detail_author_card"`
					VideoDetailAutoPipOpt struct {
						Vid string `json:"vid"`
					} `json:"video_detail_auto_pip_opt"`
					VideoDetailNavOpt struct {
						Vid string `json:"vid"`
					} `json:"video_detail_nav_opt"`
					VideoDetailRelatedRefetch struct {
						Vid string `json:"vid"`
					} `json:"video_detail_related_refetch"`
					VideoDetailResponsiveUI struct {
						Vid string `json:"vid"`
					} `json:"video_detail_responsive_ui"`
					VideoDetailYmlUI struct {
						Vid string `json:"vid"`
					} `json:"video_detail_yml_ui"`
					VideoTopic struct {
						Vid string `json:"vid"`
					} `json:"video_topic"`
					WebCreationVmok struct {
						Vid string `json:"vid"`
					} `json:"web_creation_vmok"`
					WebPlayerRefactor struct {
						Vid string `json:"vid"`
					} `json:"web_player_refactor"`
					WebappCollectionProfile struct {
						Vid string `json:"vid"`
					} `json:"webapp_collection_profile"`
					WebappCreatorJustWatched struct {
						Vid string `json:"vid"`
					} `json:"webapp_creator_just_watched"`
					WebappCreatorPostSort struct {
						Vid string `json:"vid"`
					} `json:"webapp_creator_post_sort"`
					WebappExploreCategory struct {
						Vid string `json:"vid"`
					} `json:"webapp_explore_category"`
					WebappExploreNavOrder struct {
						Vid string `json:"vid"`
					} `json:"webapp_explore_nav_order"`
					WebappExploreVideoInfo struct {
						Vid string `json:"vid"`
					} `json:"webapp_explore_video_info"`
					WebappJotaiDetail struct {
						Vid string `json:"vid"`
					} `json:"webapp_jotai_detail"`
					WebappOdinIDFeReverse struct {
						Vid string `json:"vid"`
					} `json:"webapp_odin_id_fe_reverse"`
					WebappPreviewCover struct {
						Vid string `json:"vid"`
					} `json:"webapp_preview_cover"`
					WebappRecommendLanguage struct {
						Vid string `json:"vid"`
					} `json:"webapp_recommend_language"`
					WebappRepostAction struct {
						Vid string `json:"vid"`
					} `json:"webapp_repost_action"`
					WebappRepostLabel struct {
						Vid string `json:"vid"`
					} `json:"webapp_repost_label"`
					WebappRepostNotice struct {
						Vid string `json:"vid"`
					} `json:"webapp_repost_notice"`
					WebappRepostTab struct {
						Vid string `json:"vid"`
					} `json:"webapp_repost_tab"`
					WebappSeoPhotomodeUserExp struct {
						Vid string `json:"vid"`
					} `json:"webapp_seo_photomode_user_exp"`
					WebappVideoDetailPageRelatedMask struct {
						Vid string `json:"vid"`
					} `json:"webapp_video_detail_page_related_mask"`
					Webcast struct {
						EnableEcLcc                 bool `json:"enable_ec_lcc"`
						LikeIconOptimize            bool `json:"like_icon_optimize"`
						LiveRestrictedMode          bool `json:"live_restricted_mode"`
						WebDrawerShowExplore        bool `json:"web_drawer_show_explore"`
						WebFollowGuideStrategyGroup int  `json:"web_follow_guide_strategy_group"`
					} `json:"webcast"`
					YmlUIOptimize struct {
						Vid string `json:"vid"`
					} `json:"yml_ui_optimize"`
				} `json:"parameters"`
				AbTestApp struct {
					Parameters struct {
						Tiktok struct {
							LongVideoPopupDisplayOptimization bool `json:"long_video_popup_display_optimization"`
						} `json:"tiktok"`
					} `json:"parameters"`
				} `json:"abTestApp"`
			} `json:"abTestVersion"`
			CsrfToken      string `json:"csrfToken"`
			UserAgent      string `json:"userAgent"`
			EncryptedWebid string `json:"encryptedWebid"`
			Host           string `json:"host"`
		} `json:"webapp.app-context"`
		WebappUserDetail struct {
			UserInfo struct {
				User struct {
					Id               string `json:"id"`
					ShortId          string `json:"shortId"`
					UniqueId         string `json:"uniqueId"`
					Nickname         string `json:"nickname"`
					AvatarLarger     string `json:"avatarLarger"`
					AvatarMedium     string `json:"avatarMedium"`
					AvatarThumb      string `json:"avatarThumb"`
					Signature        string `json:"signature"`
					CreateTime       int    `json:"createTime"`
					Verified         bool   `json:"verified"`
					SecUid           string `json:"secUid"`
					Ftc              bool   `json:"ftc"`
					Relation         int    `json:"relation"`
					OpenFavorite     bool   `json:"openFavorite"`
					CommentSetting   int    `json:"commentSetting"`
					CommerceUserInfo struct {
						CommerceUser bool `json:"commerceUser"`
					} `json:"commerceUserInfo"`
					DuetSetting        int    `json:"duetSetting"`
					StitchSetting      int    `json:"stitchSetting"`
					PrivateAccount     bool   `json:"privateAccount"`
					Secret             bool   `json:"secret"`
					IsADVirtual        bool   `json:"isADVirtual"`
					RoomId             string `json:"roomId"`
					UniqueIdModifyTime int    `json:"uniqueIdModifyTime"`
					TtSeller           bool   `json:"ttSeller"`
					Region             string `json:"region"`
					DownloadSetting    int    `json:"downloadSetting"`
					ProfileTab         struct {
						ShowMusicTab    bool `json:"showMusicTab"`
						ShowQuestionTab bool `json:"showQuestionTab"`
						ShowPlayListTab bool `json:"showPlayListTab"`
					} `json:"profileTab"`
					FollowingVisibility    int           `json:"followingVisibility"`
					RecommendReason        string        `json:"recommendReason"`
					NowInvitationCardUrl   string        `json:"nowInvitationCardUrl"`
					NickNameModifyTime     int           `json:"nickNameModifyTime"`
					IsEmbedBanned          bool          `json:"isEmbedBanned"`
					CanExpPlaylist         bool          `json:"canExpPlaylist"`
					ProfileEmbedPermission int           `json:"profileEmbedPermission"`
					Language               string        `json:"language"`
					EventList              []interface{} `json:"eventList"`
					SuggestAccountBind     bool          `json:"suggestAccountBind"`
					IsOrganization         int           `json:"isOrganization"`
				} `json:"user"`
				Stats struct {
					FollowerCount  int `json:"followerCount"`  // 粉丝数
					FollowingCount int `json:"followingCount"` // 关注数
					Heart          int `json:"heart"`          // 点赞
					HeartCount     int `json:"heartCount"`     // 点赞
					VideoCount     int `json:"videoCount"`     // 视频个数
					DiggCount      int `json:"diggCount"`      // 点赞
					FriendCount    int `json:"friendCount"`    // 朋友数
				} `json:"stats"`
				ItemList []interface{} `json:"itemList"`
			} `json:"userInfo"`
			ShareMeta struct {
				Title string `json:"title"`
				Desc  string `json:"desc"`
			} `json:"shareMeta"`
			StatusCode    int  `json:"statusCode"`
			StatusAndCode int  `json:"status_code"`
			NeedFix       bool `json:"needFix"`
		} `json:"webapp.user-detail"`
	} `json:"__DEFAULT_SCOPE__"`
}
