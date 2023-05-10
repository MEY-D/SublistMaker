# SublistMaker
Tool For Make A Unique Subdomain List For Dns Brute Force

# Usage
Usage: sublistmaker [OPTIONS]

Options:

  -l    wordlist path
  
  -o    output path
  
  -s    save the result in $HOME/database/sublist.txt
  
  
  ```
  subfinder -all -dL ListOfDomain.txt >> Domains.txt
  sublistmaker -l Domains.txt -s
  or
  sublistmaker -l Domains.txt
  ```
  So If You Don't Use The -s Flag The Result Just Printed
  
  And If Use it, It Will Save in $HOME/database/sublist.txt
  
  So Before Using -s Flag Please Make The database And sublist.txt
