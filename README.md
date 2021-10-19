# ReCo

ReCo is a simple Go CLI built with [Cobra](https://github.com/spf13/cobra) that will be capable of doing the followings:
</br>
</br>
Read files and calculates:
 - word count,
 - number of letters,
 - number of vowels/consonants,
 - the number of punctuation marks/spaces.
 
The name ReCo is an abbreviation for read and count.

If you want to build new command-line interfaces, here is the [blog](https://medium.com/@3n0ugh/creating-cli-in-go-with-cobra-d6f83dcbab1c).

# Instructions

- Clone the repository.
```bash
  git clone https://github.com/3n0ugh/ReCo.git
```
- Then, tidy up the packages.
```bash
  go mod tidy
```
- Lastly, build the ReCo.
```bash
  go install
```
- Now, you can use it.
```
  # Sample
  ReCo count -w -v -c -l -s -d -p -t test.txt
  
  # To see all flags
  ReCo count -h
```
