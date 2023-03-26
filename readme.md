# Email Verifier

This program checks for the presence of various DNS records associated with a given domain, to determine if it has proper email setup. Specifically, it checks for the following records: MX, SPF, and DMARC.

## Usage

To use this program, run the following command in a terminal:

go run main.go

Once the program is running, it will prompt you to enter a domain name. After entering the domain name, the program will check for the presence of the required DNS records and output the results.

The program will continue to prompt you for domain names until you terminate the program.

## Dependencies

This program uses the following Go packages:

- `bufio` for reading input from the command line
- `log` for logging errors
- `net` for DNS lookups
- `os` for reading input from the command line
- `strings` for string manipulation

## Functionality

The program defines the `checkDomain` function, which takes a domain name as a parameter and checks for the presence of the required DNS records. Specifically, it uses the `net` package to perform MX, SPF, and DMARC lookups.

If an error occurs during any of the lookups, the program logs the error and moves on to the next lookup.

Once all the lookups are complete, the program outputs the results to the command line. The output consists of the following fields:

- `domain` - the domain name that was checked
- `hasMX` - a boolean indicating whether the domain has any MX records
- `hasSPF` - a boolean indicating whether the domain has an SPF record
- `sprRecord` - the value of the SPF record, if present
- `hasDMARC` - a boolean indicating whether the domain has a DMARC record
- `dmarcRecord` - the value of the DMARC record, if present

The `main` function simply prompts the user for input and calls the `checkDomain` function for each domain entered.
