#Prerequisites

    Go 1.16 or later
    (Optional) Xsel for Linux, which can be installed with sudo apt-get install xsel

#How to use

    Clone this repository to your local machine
    Navigate to the repository directory in your terminal
    Run the command go build to build the application
    Run the command ./random_password_generator to generate a password
    The generated password will be copied to the clipboard and saved to password.txt

#Customization

The password generator can be customized by adjusting the following variables in the main() function:

    passLength: Length of the generated password
    minSpecial: Minimum number of special characters in the password
    minUpper: Minimum number of uppercase letters in the password
    minNumbers: Minimum number of numbers in the password
    minLower: Minimum number of lowercase letters in the password

#Troubleshooting

If you encounter any issues with copying the generated password to the clipboard, you may need to install Xsel for Linux using the command sudo apt-get install xsel.