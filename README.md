# PyBadge Workshop

## What you need

    - Adafruit's PyBadge
    - Personal computer with Go 1.13+ and TinyGo installed, and a serial port.

## Installation

### Go 1.13

If somehow you have not installed Go 1.13 on your computer already, you can download it here:

https://golang.org/dl/

Now you are ready to install TinyGo.

### TinyGo

Follow the instructions here for your operating system:

https://tinygo.org/getting-started/

Once you have finished installing TinyGo itself, install the drivers needed for the 

### TinyGo drivers

To install the various drivers and other code dependencies run these commands:

```
go get -u tinygo.org/x/drivers
go get -u github.com/conejoninja/tinydraw
go get -u github.com/conejoninja/tinyfont
```

### Bossa

In order to "flash", meaning to move the binary code from your computer to the PyBadge, you must install the "bossac" command line utility which is part of BOSSA.

#### Linux

On Linux, install from source:

```
sudo apt install \
  libwxgtk3.0-dev \
  libreadline-dev
git clone https://github.com/shumatech/BOSSA.git
cd BOSSA
make
```

Install Clang-8:

```
sudo apt install clang-8
```

#### macOS

On macOS, you can install it via `brew cask install bossa` or download the installer from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-1.9.1.dmg

One you have downloaded it, double click on the .dmg file to perform the installation.

#### Windows

- You must install the "BOSSA" flashing utility first. You can download it from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-x64-1.9.1.msi

- During the installation, you should choose to put it into `c:\Program Files`

- After the installation, you must add it to your PATH:

    ```shell
    set PATH=%PATH%;"c:\Program Files\BOSSA";
    ```

- Test that you have installed "BOSSA" correctly by running this command:

    ```shell
    bossac --help
    ```

## Connecting the PyBadge to your computer

![Adafruit PyBadge](./images/pybadge_hello.jpg)

Plug the PyBadge into your computer using a USB cable. There may be one provided in your starter kit.

## Running the code

The TinyGo programs will run directly on the PyBadge's microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your PyBadge.
- The program executes from the PyBadge. You can disconnect the PyBadge from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## Code

### IMPORTANT! macOS-specific required steps (Linux users can ignore)
The Linux and macOS device subsystems have subtle differences. **In order to talk to the PyBadge board on a Mac**, you need to discover how macOS system has named it. Plug it in and follow these commands:

```
ls /dev | grep usb
```

Should produce entries like (you may have different numbers after
`usbmodem`):

```
/dev/cu.usbmodem141201
/dev/tty.usbmodem141201
```

Set the fully qualified path for the `tty` entry into an environment variable. We'll use this below in the macOS version of the steps.

```sh
export PYBADGE_DEV_PATH=/dev/{YOUR TTY USBMODEM ID}
```


### step0.go - Built-in LED

This tests that you can compile and flash your PyBadge with TinyGo code, by blinking the built-in LED (it's on the back).

![Adafruit PyBadge bootloader](./images/pybadge_bootloader.jpg)

- Click on the "RST" button two times to put the PyBadge into bootloader mode so you can load your own code onto it. The front neopixels will lit green and the screen will show a message to indicate that the PyBadge is ready to receive your code.

Run the following command to compile your code, and flash it onto the PyBadge:

#### Linux
```
tinygo flash -target pybadge ./step0/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step0/main.go
```      

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step0/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

Once the PyBadge is flashed correctly, the built-in LED labeled "D13" (on the back) should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### step1.go - Built-in LED, START Button

Run the code.

#### Linux
```
tinygo flash -target pybadge ./step1/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step1/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step1/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

When you press the START button, the built-in LED should turn on.

### step2.go - Neopixels

Run the code.

#### Linux
```
tinygo flash -target pybadge ./step2/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step2/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step2/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The 5 neopixels should light up green and red alternatively.

### step3.go - Neopixels, Buttons

Run the code.

#### Linux
```
tinygo flash -target pybadge ./step3/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step3/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step3/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The 5 neopixels should light up in different colors depending on which button you press.

### step4.go - Light sensor, Neopixels

Run the code.

#### Linux
```
tinygo flash -target pybadge ./step4/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step4/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step4/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The neopixels will should lit green, they will turn red if not enough light is reaching the sensor A7 (on top of display). If there's no enough light in the room, try using your smartphone's torch to illuminate it a little bit.

### step5.go - Display

![PyBadge](./images/pybadge_hello.jpg)

Run the code.

#### Linux
```
tinygo flash -target pybadge ./step5/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step5/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step5/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The message "Hello Gophers!" should appear on the display.

### step6.go - Display, Buttons

![PyBadge](./images/pybadge_display_buttons.jpg)

Run the code.

#### Linux

```
tinygo flash -target pybadge ./step6/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step6/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step6/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The display will show some blue circles. When a button is pressed a ring will be shown around its corresponding circle.

### step7.go - Display, Accelerometer

![PyBadge](./images/pybadge_accel.jpg)

Run the code.

#### Linux

```
tinygo flash -target pybadge ./step7/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step7/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step7/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

The display will show a bar for each X,Y,Z axis. Move the Pybadge to see it in action.

### step8.go - Buzzer, Buttons

Run the code.

#### Linux

```
tinygo flash -target pybadge ./step8/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./step8/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./step8/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

Press the buttons and create your melody.

### Snake Game

![PyBadge](./images/pybadge_snake.jpg)

Run the code.

#### Linux

```
tinygo flash -target pybadge ./snake/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./snake/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./snake/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

Play the famous Snake game on the pybadge.


### My Name Is

![PyBadge](./images/pybadge_mynameis.jpg)

Run the code.

#### Linux

```
tinygo flash -target pybadge ./mynameis/main.go
```

#### macOS
```
tinygo flash -target pybadge -port=$PYBADGE_DEV_PATH ./mynameis/main.go
```

#### Windows
```
tinygo build -o arcade.uf2 -target pybadge -port=$PYBADGE_DEV_PATH ./mynameis/main.go
```                                            
And then move the arcade.uf2 file to the PYBADGEBOOT drive.

Configure your name and use the awesome TinyGo-powered badge!