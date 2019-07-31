# Benchmarking

At the moment a quick check can be made with curl.

```
$ curl -w "@utils/benchmarking/curl-format.txt" -o /dev/null -s http://localhost:1718

    time_namelookup:  0.001621
       time_connect:  0.001781
    time_appconnect:  0.000000
   time_pretransfer:  0.001864
      time_redirect:  0.000000
 time_starttransfer:  0.003671
                    ----------
         time_total:  0.003760
```
