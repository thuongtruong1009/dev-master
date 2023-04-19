# I. Execute command

-   run: go <file_name.go>

-   build: go build -o <prod_path/prod_name.exe> <main_path/main_name.go>

# II. Theory

### goroutine

-   Have send must have receive
-   Send done must be close

### context

-   cancel context: cancel all goroutine when one goroutine is done or cancel
-   timeout context: cancel all goroutine when timeout
-   value context: pass value to all goroutine

### select

-   select no dag nghe 2 tg chanel. tg nao co data dau tien thi no se chay block code cua tg do. Neu 2 tg cung co = nhau thi no se random

-   nen select thuong dung de tranh goroutine leak = cach dat timer

-   khi het time quy dinh ma cai chanel do chua tra data ve thi tat dien goroutine

-   1 channel se tu dong dc gui data vao khi het time quy dinh nen la select no se chon block code cua timer thay vi block code cua tg chanel dang can xu ly
    r trog block code cua timer minh cho out cmnl la xog

### others

-   lazy function: must have time to init and assign value to variable -> needed time.Sleep
-   init function: init() is called and initialize variables before main() -> not need time.Sleep

### design pattern

-   builder pattern: use when we have many optional parameters

# III. Tutorial

-   https://gobyexample.com/variables

-   https://nordiccoder.com/blog/golang/

-   https://youtu.be/h2RdcrMLQAo

-   https://www.youtube.com/watch?v=bM6N-vgPlyQ&list=PLy-NDN51bIDVUNrl5KpfdHqkHfpFEFvWW

-   https://www.youtube.com/watch?v=VkGQFFl66X4&pp=ugMICgJ2aRABGAE%3D

-   https://www.youtube.com/watch?v=IS9Erqkx-ks&list=PLGRs9mrW62D1I4SF1p03zNn0x2WSSFhOA

-   https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE

# IV. Projects
