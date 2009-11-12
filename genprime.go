package main

import ("fmt"; "os"; "math"; "strconv")

func getTime() float64 {
    sec, nsec, err := os.Time();
    var time float64 = 0;
    if err == nil {
        time = float64(sec) + (float64(nsec) * math.Pow10(-9));
    }
    return time;
}

func isprime(x int) int {
    if x < 2 {
        return 0;
    }
    if x < 4 {
        return 1;
    }
    if x == 5 {
        return 1;
    }
    if x % 2 == 0 {
        return 0;
    }
    if x % 5 == 0 {
        return 0;
    }
    if (x + 1) % 6 != 0 {
        if (x - 1) % 6 != 0 {
            return 0;
        }
    }
    match := 0;
    lim := math.Sqrt(float64(x)) + 1.0;
    for y := 3; float64(y) < lim; y += 2 {
        if x % y == 0 {
            match = 1;
            break;
        }
    }
    if match == 1 {
        return 0;
    }
    return 1;
}

func genprime(max int) int {
    current := 1;
    for count := 0; count < max; current++ {
        if isprime(current) == 1 {
            count++;
        }
    }
    return (current - 1);
}

func main() {
    var err os.Error = nil;
    var start int = 0;
    var stop int = 0;
    if len(os.Args) > 1 {
        start, err = strconv.Atoi(os.Args[1]);
        if err != nil {
            os.Exit(1);
        }
    }
    if len(os.Args) > 2 {
        stop, err = strconv.Atoi(os.Args[2]);
        if err != nil {
            os.Exit(2);
        }
    }

    for x := start; x <= stop; x += start {
        begin := getTime();
        last := genprime(x);
        end := getTime();
        var duration float64 = (end - begin);
        fmt.Printf("Found %8d primes in %10f seconds (last was %8d)\n", x, duration, last)
    }
}
