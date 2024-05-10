package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// BeefSummary เก็บข้อมูลเกี่ยวกับจำนวนเนื้อแต่ละชนิดที่เกี่ยวข้องกับ beef
type BeefSummary map[string]int

// isBeef ตรวจสอบว่าคำเป็นชื่อของเนื้อ "beef" หรือไม่
func isBeef(word string) bool {
	beefNames := map[string]bool{
		"bacon":    true,
		"filet":    true,
		"mignon":   true,
		"ribeye":   true,
		"flank":    true,
		"t-bone":   true,
		"pastrami": true,
		"pork":     true,
		"meatloaf": true,
		"jowl":     true,
		"bresaola": true,
	}

	// ตรวจสอบว่าคำอยู่ใน map หรือไม่
	_, found := beefNames[word]
	return found
}

// CountBeef นับจำนวนเนื้อแต่ละชนิดที่เกี่ยวข้องกับ beef
func CountBeef(text string) BeefSummary {
	// สร้าง map เพื่อเก็บจำนวนเนื้อแต่ละชนิด
	counts := make(BeefSummary)

	// แยกคำออกจากข้อความ
	words := strings.Fields(text)

	// นับจำนวนเนื้อแต่ละชนิดที่เกี่ยวข้องกับ beef
	for _, word := range words {
		if isBeef(word) {
			counts[word]++
		}
	}

	return counts
}

func BeefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	// ดึงข้อมูลจาก URL
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// อ่านข้อมูล response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	// แปลงข้อมูลเป็น string
	text := string(body)

	// นับจำนวนเนื้อแต่ละชนิดที่เกี่ยวข้องกับ beef
	summary := CountBeef(text)

	// แปลงข้อมูลเป็น JSON
	jsonData, err := json.Marshal(summary)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// ตอบกลับด้วย JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	// กำหนด route /beef/summary
	http.HandleFunc("/beef/summary", BeefSummaryHandler)

	// เริ่มต้นเซิร์ฟเวอร์
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
