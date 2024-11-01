<p align="center"><img src="./images/image_3.png" width="50%"/></p><br/>
<p align="center">Various detection mechanisms for anti-debugging and anti-virtual machine (VM) environments.</p>
<p align="center">made in golang.</p>

</p><br/><br/>

# Anti VM
## Triage
| Platform                      | Background Detection | Serials Detection | Date/time  | Screenshot                       |
|-------------------------------|----------------------|-------------------|------------|----------------------------------|
| Windows 10 (2004) x64         | **Triggered**        | **Triggered**     | 26/10/2024 | ![](./images/img_2.png)        |
| Windows 10 (LTSC 2021) x64    | **Triggered**        | **Triggered**     | 26/10/2024 | ![](./images/img_1.png)        |
| Windows 11 (21H2) x64         | **Triggered**        | **Triggered**     | 26/10/2024 | ![](./images/img.png)          |
## VMWare
| Platform              | Diskdrive Detection | BIOS Detection | Model Detection | Date/time  | Screenshot              |
|-----------------------|---------------------|----------------|-----------------|------------|-------------------------|
| Windows 11 (24H2) x64 | **Triggered**       | **Triggered**  | **Triggered**   | 27/10/2024 | ![](./images/img_3.png) |
| Windows 10 (22H2) x64 | **Triggered**       | **Triggered**  | **Triggered**   | 27/10/2024 | ![](./images/img_5.png) |
## Windows SandBox
| Platform              | Username Detection | Date/time    | Screenshot              |
|-----------------------|--------------------|--------------|-------------------------|
| Windows 11 (24H2) x64 | **Triggered**      | 27/10/2024   | ![](./images/img_4.png) |
## VirtualBox
| Platform              | Motherboard Detection | Discdrive Detection | BIOS Detection | Date/time    | Screenshot              |
|-----------------------|-----------------------|---------------------|----------------|--------------|-------------------------|
| Windows 11 (24H2) x64 | **Triggered**         | **Triggered**       | **Triggered**  | 27/10/2024   | ![](./images/img_6.png) |

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/9dl/goware.git
    ```
2. Navigate to the project directory:
    ```sh
    cd goware
    ```
3. Build the project:
    ```sh
    go build
    ```

## Usage

Run the compiled binary:
```sh
./goware
