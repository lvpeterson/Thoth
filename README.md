# Thoth - Hash Processing Utility

Thoth is a powerful command-line utility designed for processing and managing various types of hash formats. It provides functionality for sanitizing hash files, processing cracked hashes, and handling different hash algorithms commonly found in penetration testing.

## Features

- **Hash Sanitization**: Sanitize hash files for various formats including NTLM, LM, NetNTLMv1/v2, and Kerberos
- **Desanitization**: Combine cracked results with original hashes
- **LM Hash Utility**: Special handling for LM hash processing and wordlist generation
- **Support for Multiple Hash Types**:
    - NTLM (Mode 1000)
    - LM (Mode 3000)
    - NetNTLMv1 (Mode 5500)
    - NetNTLMv2 (Mode 5600)
    - Kerberos TGS (Mode 13100)
    - Kerberos AS-REP (Mode 18200)

## Installation

`go install github.com/lvpeterson/Thoth@latest`

## To Do
Desanitization support for the following modes:
- 5500
- 5600
- 13100
- 18200