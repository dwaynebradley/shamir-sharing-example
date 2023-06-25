# Shamir Secret Sharing Example

Sharing secret credentials among colleagues is simple. I can just store them in BitWarden or 1Password or whatever password manager I want, put them in a group that my colleagues have access to, and voila! The secret is shared.

What if I want to do better than that?

Suppose I want to have a secret shared between four people. To retrieve the secret, however, two of those four must provide information. Or suppose I have two groups of people, one being the C-suite and the other being a group of engineers. Is there a way to craft a secret so that any one of the executives can retrieve it, but it takes two engineers to do so?

There is.

From [Wikipdia](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing):

> Shamir's secret sharing (SSS) is an efficient secret sharing algorithm for distributing private information
(the "secret") among a group so that the secret cannot be revealed unless a quorum of the group acts together to 
pool their knowledge. To achieve this, the secret is mathematically divided into parts (the "shares") from which 
the secret can be reassembled only when a sufficient number of shares are combined. SSS has the property of 
information-theoretic security, meaning that even if an attacker steals some shares, it is impossible for the 
attacker to reconstruct the secret unless they have stolen the quorum number of shares. 

[Hashicorp Vault]() implements Shamir's Secret Sharing. In this example their library is leveraged.

## Building the code

```
$ go get
$ go build -o bin/shamir
```

## Running the code

```
$ bin/shamir
Secret: p455w0rdhunt3r2
Part 0: WDiR3HaLzp5qgaVfLfqutQ==
Part 1: 4x9stK+FQ1IVcgYlvHPJhA==
Part 2: Y+FE3nR0Ak30AWbscwdzGQ==
Part 3: isqNdercvWeMUdjKuRi1yw==
Part 4: eIOA6FClgg3bqD9YlpfjIQ==
[1] Enter any unique part of the original 5 parts: WDiR3HaLzp5qgaVfLfqutQ==
[2] Enter any unique part of the original 5 parts: 4x9stK+FQ1IVcgYlvHPJhA==
[3] Enter any unique part of the original 5 parts: eIOA6FClgg3bqD9YlpfjIQ==
Retreived secret: p455w0rdhunt3r2
```