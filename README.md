# Random Password Generator

This is a command-line tool that generates a random password, copies it to the clipboard and stores it in a file named password.txt.

## Prerequisites

    Go 1.16 or later
    (Optional) Xsel for Linux, which can be installed with sudo apt-get install xsel

## How to use

    Clone this repository to your local machine
    Navigate to the repository directory in your terminal
    Run the command go build to build the application
    Run the command ./random_password_generator to generate a password
    The generated password will be copied to the clipboard and saved to password.txt

## Customization

The password generator can be customized by adjusting the following variables in the main() function:

    passLength: Length of the generated password
    minSpecial: Minimum number of special characters in the password
    minUpper: Minimum number of uppercase letters in the password
    minNumbers: Minimum number of numbers in the password
    minLower: Minimum number of lowercase letters in the password

## Dependencies

 This program does not require any external dependencies. All necessary packages are part of the Go standard library.

## Supported Operating Systems

 This program is supported on Windows, macOS, and most Linux distributions. However, the clipboard functionality may not work on all Linux distributions. If the clipboard functionality does not work on your Linux distribution, the generated password can still be found in the password.txt file.

## Troubleshooting

If you encounter any issues with copying the generated password to the clipboard, you may need to install Xsel for Linux using the command sudo apt-get install xsel.
