package main

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	url := "https://www.douyin.com/aweme/v1/web/aweme/post/?device_platform=webapp&aid=6383&channel=channel_pc_web&sec_user_id=MS4wLjABAAAAdFc69YowYQxY2YlTWU0riFYRkyWRWzMqwkksjVknNhZcs11dsPxilGKVFDEEY7k3&max_cursor=0&locate_query=false&show_live_replay_strategy=1&count=18&publish_video_strategy_type=2&pc_client_type=1&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1536&screen_height=864&browser_language=zh-CN&browser_platform=Win32&browser_name=Chrome&browser_version=115.0.0.0&browser_online=true&engine_name=Blink&engine_version=115.0.0.0&os_name=Windows&os_version=10&cpu_core_num=8&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=0&webid=7267744665397741090&msToken=KmXs9WfitWmISe0j4ud6h-5sLtm_kwPWFMg2xSoXct6Q4gTMOe9TG8Y6ey-Afylwh8TnRe8-RSuZuYNVvc089d6TlQRYR8fvA7xdo7oZe1QQWLG31Q==&X-Bogus=DFSzswVOTibANyOEt9iqdM9WX7nr"
	common.InitDB()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "ttwid=1%7CtjM-Za1wb1AGsOHtbxAuyU_YncZrBPM88ErcvKjt_1E%7C1692153682%7C3733243c15d88708d5666464a11356f1c10b76d10a5df58d45887e5dcd20eb51; strategyABtestKey=%221692153683.761%22; passport_csrf_token=fde947f6fc06db18e07bc0263ab2c515; passport_csrf_token_default=fde947f6fc06db18e07bc0263ab2c515; s_v_web_id=verify_lld4m37p_Jq4E0MhH_tYcb_4LLS_8N0c_k5Nov1ZHhACe; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Afalse%2C%22volume%22%3A0.5%7D; xgplayer_user_id=216292960654; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtY2xpZW50LWNzciI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tXHJcbk1JSUJEakNCdFFJQkFEQW5NUXN3Q1FZRFZRUUdFd0pEVGpFWU1CWUdBMVVFQXd3UFltUmZkR2xqYTJWMFgyZDFcclxuWVhKa01Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRU00SE9kT3pTWUpybjJScC9wTENKeG1vclxyXG5Jdkl6SHdkM2todUc2dXN4b2RvK3lnVm5sQTJqUlQweWdiMGhyellpdHF5SlZyeWRucjhuQTBrTFA1d0tVNkFzXHJcbk1Db0dDU3FHU0liM0RRRUpEakVkTUJzd0dRWURWUjBSQkJJd0VJSU9kM2QzTG1SdmRYbHBiaTVqYjIwd0NnWUlcclxuS29aSXpqMEVBd0lEU0FBd1JRSWdVTmd1RzVqUTA5S3M2VzloRWJVWmMzNUs0bFhVa1QvaEkyNFpNbnFXbXZnQ1xyXG5JUUROTjhKcnNmRUViYWVrS0lncFhpL0lBZmRLU0o3ZnhwMm8zNEM5aHJFUXlBPT1cclxuLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tXHJcbiJ9; ttcid=8642ee4a3dbb4a8f9ef0d02af1c393e012; VIDEO_FILTER_MEMO_SELECT=%7B%22expireTime%22%3A1692761559770%2C%22type%22%3Anull%7D; SEARCH_RESULT_LIST_TYPE=%22single%22; FORCE_LOGIN=%7B%22videoConsumedRemainSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; download_guide=%223%2F20230816%2F1%22; passport_assist_user=CkAJGrhDi6aj3XsA8m4dtsv1o8XVSRfP4G7EI7BGiUfD1PlmTZ2CdRIKi5C0LuNgFOgMTW2VIovaSGVVznC2YL3FGkgKPFVie_vAG0Y_8Io2hmR9cTy_Q13DQ1CPfV0kitoLAq5EXfmZsHPaFYE6WvvJV8KoQN9SS7CNYKESSErSzxC0qrkNGImv1lQiAQOyk3ZM; n_mh=YA4qYbyO8tqygaeMONEslSUzBaIGNp_WX7Z8XDT4JxA; sso_uid_tt=752be07fd738f35e2d3b42e22d72e5cb; sso_uid_tt_ss=752be07fd738f35e2d3b42e22d72e5cb; toutiao_sso_user=044048209791e5e7c26fd2285612452d; toutiao_sso_user_ss=044048209791e5e7c26fd2285612452d; sid_ucp_sso_v1=1.0.0-KDFhN2YwZDg1ODVkZjgzY2U4NWIwZmE4YWM5Njg3M2ZlNTUyMTE3N2MKHwisi8Cv1YymARCbh_GmBhjvMSAMMPzv1YoGOAZA9AcaAmxxIiAwNDQwNDgyMDk3OTFlNWU3YzI2ZmQyMjg1NjEyNDUyZA; ssid_ucp_sso_v1=1.0.0-KDFhN2YwZDg1ODVkZjgzY2U4NWIwZmE4YWM5Njg3M2ZlNTUyMTE3N2MKHwisi8Cv1YymARCbh_GmBhjvMSAMMPzv1YoGOAZA9AcaAmxxIiAwNDQwNDgyMDk3OTFlNWU3YzI2ZmQyMjg1NjEyNDUyZA; passport_auth_status=049fec03c31c65cbe69151f3b69f5b46%2C; passport_auth_status_ss=049fec03c31c65cbe69151f3b69f5b46%2C; uid_tt=d75eef814def18682a078c1f325f393e; uid_tt_ss=d75eef814def18682a078c1f325f393e; sid_tt=3ec0e3f9ac18dead4e3b85d677759757; sessionid=3ec0e3f9ac18dead4e3b85d677759757; sessionid_ss=3ec0e3f9ac18dead4e3b85d677759757; publish_badge_show_info=%220%2C0%2C0%2C1692156829960%22; __security_server_data_status=1; LOGIN_STATUS=1; store-region=cn-sd; store-region-src=uid; d_ticket=77454a331efe703a8eaf672f6e29587518620; sid_guard=3ec0e3f9ac18dead4e3b85d677759757%7C1692156866%7C5183965%7CSun%2C+15-Oct-2023+03%3A33%3A51+GMT; sid_ucp_v1=1.0.0-KDkyZmE5OGNlZmFhYTI3MzU3Y2I1N2YwNGE5Y2QzZjllN2UwOTE5NGMKGwisi8Cv1YymARDCh_GmBhjvMSAMOAZA9AdIBBoCbHEiIDNlYzBlM2Y5YWMxOGRlYWQ0ZTNiODVkNjc3NzU5NzU3; ssid_ucp_v1=1.0.0-KDkyZmE5OGNlZmFhYTI3MzU3Y2I1N2YwNGE5Y2QzZjllN2UwOTE5NGMKGwisi8Cv1YymARDCh_GmBhjvMSAMOAZA9AdIBBoCbHEiIDNlYzBlM2Y5YWMxOGRlYWQ0ZTNiODVkNjc3NzU5NzU3; pwa2=%220%7C0%7C3%7C0%22; _bd_ticket_crypt_cookie=da4743a8e3ec01201b6fad5d2640c2de; odin_tt=b25fcb8ddb236937f2856948533639baabddc22a15c05eda08d57f2ce96e608f4dfb42ad6530fcaa13281cf32c73c64915b3c5d1e9cc953e1aac20769fe00ac6; __ac_nonce=064dca2b100badf84a94; __ac_signature=_02B4Z6wo00f0164E7tgAAIDDLgYUmUNl7meuJOpAAI9mBv9bg6Ylt-YMYVBwUjZ1oUki9StWf7H4HwuYFGKdKhSw.xWqaFx4mQxxfX26Df6jrE.VtwqOEAdJHIeh7jSwWu5hmWmFSRBZajHA2c; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A1536%2C%5C%22screen_height%5C%22%3A864%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A8%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A0%7D%22; msToken=9u5HwhlxkNxLosgEjOwGvRZIjFpZOgfUQCz04cKocyxFyuTU0NFLPcxEgDO2t5gHHeALoddahoq-W-0-WDpo3Ixo18CNv24DTDTGYr2oXWCtDPJb179o4WqCOgA=; msToken=KmXs9WfitWmISe0j4ud6h-5sLtm_kwPWFMg2xSoXct6Q4gTMOe9TG8Y6ey-Afylwh8TnRe8-RSuZuYNVvc089d6TlQRYR8fvA7xdo7oZe1QQWLG31Q==; tt_scid=UQ4SR5D1sShoNrodDqvL2n-85mhrDgHW7n5ANL-BGJwABSBdCj5SaL157LjWEpwndc74; passport_fe_beating_status=false; home_can_add_dy_2_desktop=%221%22")
	req.Header.Add("Referer", "https://www.douyin.com/user/MS4wLjABAAAAdFc69YowYQxY2YlTWU0riFYRkyWRWzMqwkksjVknNhZcs11dsPxilGKVFDEEY7k3")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	awemeList := result["aweme_list"].([]interface{})
	var id int64 = 1
	for _, item := range awemeList {
		video := item.(map[string]interface{})["video"].(map[string]interface{})["play_addr"].(map[string]interface{})["url_list"].([]interface{})[0].(string)
		image := item.(map[string]interface{})["video"].(map[string]interface{})["cover"].(map[string]interface{})["url_list"].([]interface{})[0].(string)
		title := item.(map[string]interface{})["desc"].(string)
		db := common.GetDB()
		var userInfoTable model.UserInfoTable
		db.Take(&userInfoTable, 1)
		playUrl := "http://192.168.37.1:8080/" + "static/" + strconv.FormatInt(id, 10) + ".mp4"
		videoTable := model.VideoTable{
			Id:              id,
			UserInfoTableId: 1,
			User:            userInfoTable,
			PlayUrl:         playUrl,
			CoverUrl:        image + ".jpg",
			Title:           title,
		}
		id += 1
		db.Create(&videoTable)
		videoResp, err := http.Get(video)
		if err != nil {
			fmt.Println("Error downloading video:", err)
			return
		}
		defer videoResp.Body.Close()

		imgResp, err := http.Get(image)
		if err != nil {
			fmt.Println("Error downloading image:", err)
			return
		}
		defer imgResp.Body.Close()

		videoData, err := ioutil.ReadAll(videoResp.Body)
		if err != nil {
			fmt.Println("Error reading video content:", err)
			return
		}

		imgData, err := ioutil.ReadAll(imgResp.Body)
		if err != nil {
			fmt.Println("Error reading image content:", err)
			return
		}

		videoFilename := fmt.Sprintf("%d.mp4", id-1)
		imgFilename := fmt.Sprintf("%d%s.jpg", id-1, title)

		videoFile, err := os.Create(videoFilename)
		if err != nil {
			fmt.Println("Error creating video file:", err)
			return
		}
		defer videoFile.Close()

		_, err = videoFile.Write(videoData)
		if err != nil {
			fmt.Println("Error writing video content:", err)
			return
		}

		imgFile, err := os.Create(imgFilename)
		if err != nil {
			fmt.Println("Error creating image file:", err)
			return
		}
		defer imgFile.Close()

		_, err = imgFile.Write(imgData)
		if err != nil {
			fmt.Println("Error writing image content:", err)
			return
		}
	}
}
