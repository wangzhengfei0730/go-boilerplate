# Gracefully Shutdown

Use semaphore (sync.WaitGroup in Go) to implement gracefully shutdown feature. Normal program will print 10 times in the console, once use use Ctrl + C to terminate the program, it will wait for 3 more print and then exit.
