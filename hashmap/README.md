# hash generator

this is just a simple experiement

## how does it works?
1. we define a string
2. we define a hash as uint32 with start value zero
3. we iterate over the string length
4. per iteration we calculate recaculate the hash
    - with the prime number 31. a prime number avoids periodic pattern. has no small divsior. equally distributes values in the module space
     - `hash = (((0*31 + 'a')*31 + 'b')*31 + 'c')`
     - `a'*31² + 'b'*31¹ + 'c'*31`
    - we add the current byte as u32int to the calculation
5. we return the hashed string
