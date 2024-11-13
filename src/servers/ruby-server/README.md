## Ruby Server Setup

This guide provides instructions to set up and run a simple Ruby server using Sinatra. This server allows you to add, update, delete, and fetch users.

### Prerequisites

Ensure you have **Ruby** and **Bundler** installed on your machine.


#### Windows

1. **Install Ruby**:
   - Visit [rubyinstaller.org](https://rubyinstaller.org/) to download and install Ruby. 
   - Follow the installer steps, making sure to check the option to add Ruby to your PATH.

2. **Install Bundler**:
   - Open Command Prompt and run:
     ```bash
     gem install bundler
     ```

#### macOS

1. **Install Ruby**:
   - Ruby is pre-installed on most macOS versions. To check the version or install updates, you can use [Homebrew](https://brew.sh/) (if necessary):
     ```bash
     brew install ruby
     ```
   - Verify your installation with:
     ```bash
     ruby -v
     ```

2. **Install Bundler**:
   - Open Terminal and run:
     ```bash
     gem install bundler
     ```

### Installation

1. Navigate to the root directory of the project.

2. Create a Gemfile with the following dependencies:


    ```ruby
    source 'https://rubygems.org'

    gem 'sinatra'
    gem 'json'

3. Install the required gems:

    ```bash
    bundle install
    ```

### Running the Server

Use the following command to start the Ruby server:

```bash
ruby src/servers/ruby-server/server.rb
```

Make sure the server runs on **port 5003**.

### To Stop the Server

Press `Ctrl + C` in the terminal to stop the server.