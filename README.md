# Tubes2_AFaTar
Implementation of IDS dan BFS Algorithm in WikiRace

<!-- 
minimal berisi:
i. Penjelasan singkat algoritma IDS dan BFS yang diimplementasikan
ii. Requirement program dan instalasi tertentu bila ada
iii. Command atau langkah-langkah dalam meng-compile atau build program
iv. Author (identitas pembuat) 
-->

# 🏆 WikiRace

WikiRace is a competitive online game where players navigate through Wikipedia by clicking hyperlinks to reach a specific target page from a given starting page in the shortest number of clicks possible. For more information:

- [Project Specification](https://docs.google.com/document/d/1h6WY_NxfCBPrKkS84Crm2qAhrRA8DatL/edit)

## IDS and BFS Algorithm
### BFS
Algoritma Breadth-First Search (BFS) atau algoritma pencarian melebar merupakan salah satu algoritma untuk melakukan traversal pada graf. Algoritma ini menelusuri dari sebuah simpul akar pada graf lalu mengunjungi semua tetangga dari simpul tersebut. Algoritma BFS memerlukan sebuah antrian (queue) untuk menyimpan simpul yang telah dikunjungi.
### IDS
Iterative-Deepening Search (IDS) merupakan variasi dari Depth-First Search (DFS) yang bertujuan untuk menemukan solusi dalam graf dengan menggunakan strategi pencarian secara berulang dengan kedalaman pencarian yang bertambah. Meskipun konsep dasar dari IDS mirip dengan DFS, IDS membatasi kedalaman pencarian pada setiap iterasi untuk mengatasi beberapa kelemahan DFS, seperti ketidakmampuan untuk menangani graf dengan kedalaman yang tak terbatas atau terlalu dalam.

## Getting Started

### Prerequisites
1. Install Golang
2. Install npm

### Installing Dependencies 🔨

1. Clone this repository and move to the root of this project's directory

   ```
   git clone https://github.com/attaramajesta/Tubes2_AFaTar.git
   ```

2. Install dependencies

   ```
   cd src/frontend
   ```
   ```
   npm i
   ```

## How to Run 💻

1. open the web

   ```
   cd src/frontend
   ```
   ```
   npm run dev
   ```

2. open the backend in another terminal

   ```
   cd src/backend
   ```
   ```
   go run main.go
   ```

   <b>Warning if any: change permissions, etc.</b>

#### Note:

- if there's a permission pop-up, pick allow, it will not affect your device 

## Author 🪙

- [Attara Majesta Ayub](https://github.com/attaramajesta) - 13522139
- [Farrel Natha Saskoro](https://github.com/fnathas) - 13522145
- [Axel Santadi Warih](https://github.com/AxelSantadi) - 13522155
