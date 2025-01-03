# goland-lessons
GoLand lessons for beginners

## Purpose
This repository is created to help beginners to learn Go programming language using GoLand IDE.

## Prerequisites
1. Install [Go](https://golang.org/doc/install)
2. Install [GoLand](https://www.jetbrains.com/go/download/)

## How to use
1. Clone [this repository](https://github.com/ssoto/goland-lessons):
```shell
git clone <repository-url>
```
2. Open the repository in GoLand
3. Execute the code in the `main.go` file:
```shell
go run main.go
```
4. Execute a lesson using the Makefile:
```shell
make lesson-<lesson-number>
```
For example:
```shell
make run ARGS=lesson1                        
Running...
go run ./main.go lesson1
lesson1 solution has 138 elements:  2002,2009,2016,2023,2037,2044,2051,2058,2072,2079,2086,2093,2107,2114,2121,2128,2142,2149,2156,2163,2177,2184,2191,2198,2212,2219,2226,2233,2247,2254,2261,2268,2282,2289,2296,2303,2317,2324,2331,2338,2352,2359,2366,2373,2387,2394,2401,2408,2422,2429,2436,2443,2457,2464,2471,2478,2492,2499,2506,2513,2527,2534,2541,2548,2562,2569,2576,2583,2597,2604,2611,2618,2632,2639,2646,2653,2667,2674,2681,2688,2702,2709,2716,2723,2737,2744,2751,2758,2772,2779,2786,2793,2807,2814,2821,2828,2842,2849,2856,2863,2877,2884,2891,2898,2912,2919,2926,2933,2947,2954,2961,2968,2982,2989,2996,3003,3017,3024,3031,3038,3052,3059,3066,3073,3087,3094,3101,3108,3122,3129,3136,3143,3157,3164,3171,3178,3192,3199
```

You also can compile the code using the Makefile:
```shell
make build
```
The Makefile will generate a binary in the `bin/` directory. You can execute the binary using the following command:
```shell
./bin/golang-exercises lesson1
```

## Add a new lesson
This project uses cobra to manage the commands. To add a new lesson, you need to create a new command in the `cmd` directory. 
The command should be created using the following command:
```shell
cobra-cli add <lesson-name>
```
This command will create a new directory in the `cmd` directory with the lesson name. Inside the directory, you will find the `lesson-name.go` file.


## Links
- The lessons are based on Carstens [101 gonlang exercises](https://github.com/cblte/100-golang-exercises) repository.