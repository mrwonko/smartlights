EESchema Schematic File Version 4
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "Smartlights Pi HAT"
Date ""
Rev "2.5"
Comp ""
Comment1 "Not actually a HAT as it comes without software."
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Transistor_FET:IRLZ34N Q4
U 1 1 5E6648A2
P 3550 4300
F 0 "Q4" H 3754 4346 50  0000 L CNN
F 1 "IRLZ34N" H 3754 4255 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 4225 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 4300 50  0001 L CNN
	1    3550 4300
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q2
U 1 1 5E66544B
P 3550 4750
F 0 "Q2" H 3754 4796 50  0000 L CNN
F 1 "IRLZ34N" H 3754 4705 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 4675 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 4750 50  0001 L CNN
	1    3550 4750
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q3
U 1 1 5E665F41
P 3550 2500
F 0 "Q3" H 3754 2546 50  0000 L CNN
F 1 "IRLZ34N" H 3754 2455 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 2425 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 2500 50  0001 L CNN
	1    3550 2500
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q5
U 1 1 5E66772F
P 3550 2950
F 0 "Q5" H 3754 2996 50  0000 L CNN
F 1 "IRLZ34N" H 3754 2905 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 2875 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 2950 50  0001 L CNN
	1    3550 2950
	-1   0    0    1   
$EndComp
Wire Wire Line
	1100 4250 1350 4250
Wire Wire Line
	1350 4250 1350 4950
Wire Wire Line
	1350 4950 2150 4950
Wire Wire Line
	1450 4150 1100 4150
$Comp
L Connector_Generic:Conn_01x02 J3
U 1 1 5E67B518
P 1950 3250
F 0 "J3" V 1914 3062 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 3062 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 3250 50  0001 C CNN
F 3 "~" H 1950 3250 50  0001 C CNN
	1    1950 3250
	-1   0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J4
U 1 1 5E67BE71
P 1950 2800
F 0 "J4" V 1914 2612 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 2612 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 2800 50  0001 C CNN
F 3 "~" H 1950 2800 50  0001 C CNN
	1    1950 2800
	-1   0    0    -1  
$EndComp
Wire Wire Line
	9050 1650 9650 1650
$Comp
L Connector:Conn_Coaxial_Power J10
U 1 1 5E60BF44
P 6700 1150
F 0 "J10" V 6483 1100 50  0000 C CNN
F 1 "Conn_Coaxial_Power" V 6574 1100 50  0000 C CNN
F 2 "Connector_BarrelJack:BarrelJack_Horizontal" H 6700 1100 50  0001 C CNN
F 3 "~" H 6700 1100 50  0001 C CNN
	1    6700 1150
	0    1    1    0   
$EndComp
Wire Wire Line
	6800 1150 6800 1450
Connection ~ 6800 1150
Wire Wire Line
	7000 1150 6800 1150
Wire Wire Line
	7450 2350 7450 2400
Connection ~ 7450 2350
Wire Wire Line
	7650 2350 7450 2350
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5E658922
P 7650 2350
F 0 "#FLG0102" H 7650 2425 50  0001 C CNN
F 1 "PWR_FLAG" H 7650 2523 50  0000 C CNN
F 2 "" H 7650 2350 50  0001 C CNN
F 3 "~" H 7650 2350 50  0001 C CNN
	1    7650 2350
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5E6578F4
P 6800 1450
F 0 "#FLG0101" H 6800 1525 50  0001 C CNN
F 1 "PWR_FLAG" H 6800 1623 50  0000 C CNN
F 2 "" H 6800 1450 50  0001 C CNN
F 3 "~" H 6800 1450 50  0001 C CNN
	1    6800 1450
	0    -1   -1   0   
$EndComp
$Comp
L power:+12V #PWR01
U 1 1 5E65062B
P 6800 950
F 0 "#PWR01" H 6800 800 50  0001 C CNN
F 1 "+12V" H 6815 1123 50  0000 C CNN
F 2 "" H 6800 950 50  0001 C CNN
F 3 "" H 6800 950 50  0001 C CNN
	1    6800 950 
	1    0    0    -1  
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q1
U 1 1 5E60A5DF
P 3550 3400
F 0 "Q1" H 3754 3446 50  0000 L CNN
F 1 "IRLZ34N" H 3754 3355 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 3325 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 3400 50  0001 L CNN
	1    3550 3400
	-1   0    0    1   
$EndComp
$Comp
L Diode:1N5820 D1
U 1 1 5E625494
P 8250 1850
F 0 "D1" V 8204 1929 50  0000 L CNN
F 1 "1N5820" V 8295 1929 50  0000 L CNN
F 2 "Diode_THT:D_DO-201AD_P15.24mm_Horizontal" H 8250 1675 50  0001 C CNN
F 3 "http://www.vishay.com/docs/88526/1n5820.pdf" H 8250 1850 50  0001 C CNN
	1    8250 1850
	0    1    1    0   
$EndComp
Connection ~ 9050 1650
Wire Wire Line
	9050 1450 9050 1650
Wire Wire Line
	8000 1450 9050 1450
Wire Wire Line
	8250 1650 8000 1650
Connection ~ 8250 1650
Wire Wire Line
	8250 1700 8250 1650
Wire Wire Line
	8500 1650 8250 1650
Wire Wire Line
	9050 1650 9050 1700
Wire Wire Line
	8800 1650 9050 1650
Connection ~ 8250 2100
Wire Wire Line
	9050 2100 9050 2000
Wire Wire Line
	8250 2100 9050 2100
Connection ~ 7500 2100
Wire Wire Line
	8250 2100 8250 2000
Wire Wire Line
	7500 2100 8250 2100
Connection ~ 7000 2100
Wire Wire Line
	6800 2100 7000 2100
Wire Wire Line
	6800 1950 6800 2100
Connection ~ 7450 2100
Wire Wire Line
	7000 2100 7450 2100
Wire Wire Line
	7000 1650 7000 2100
Wire Wire Line
	7450 2100 7450 2350
Wire Wire Line
	7500 2100 7450 2100
Wire Wire Line
	7500 1850 7500 2100
Connection ~ 6800 1450
Wire Wire Line
	7000 1450 6800 1450
Wire Wire Line
	6800 950  6800 1150
$Comp
L power:GND #PWR02
U 1 1 5E61B63C
P 7450 2400
F 0 "#PWR02" H 7450 2150 50  0001 C CNN
F 1 "GND" H 7455 2227 50  0000 C CNN
F 2 "" H 7450 2400 50  0001 C CNN
F 3 "" H 7450 2400 50  0001 C CNN
	1    7450 2400
	1    0    0    -1  
$EndComp
$Comp
L Device:L L1
U 1 1 5E61909A
P 8650 1650
F 0 "L1" V 8469 1650 50  0000 C CNN
F 1 " 33µH" V 8560 1650 50  0000 C CNN
F 2 "Inductor_THT:L_Radial_D12.5mm_P7.00mm_Fastron_09HCP" H 8650 1650 50  0001 C CNN
F 3 "~" H 8650 1650 50  0001 C CNN
	1    8650 1650
	0    1    1    0   
$EndComp
$Comp
L Device:C C2
U 1 1 5E6170BF
P 9050 1850
F 0 "C2" H 9165 1896 50  0000 L CNN
F 1 "220µF 5V" H 9165 1805 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D6.3mm_P2.50mm" H 9088 1700 50  0001 C CNN
F 3 "~" H 9050 1850 50  0001 C CNN
	1    9050 1850
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 5E60F67B
P 6800 1800
F 0 "C1" H 6915 1846 50  0000 L CNN
F 1 "680µF 12V" H 6915 1755 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D10.0mm_P5.00mm" H 6838 1650 50  0001 C CNN
F 3 "~" H 6800 1800 50  0001 C CNN
	1    6800 1800
	1    0    0    -1  
$EndComp
$Comp
L Regulator_Switching:LM2596T-5 U1
U 1 1 5E60D4BA
P 7500 1550
F 0 "U1" H 7500 1917 50  0000 C CNN
F 1 "LM2596T-5" H 7500 1826 50  0000 C CNN
F 2 "Package_TO_SOT_SMD:TO-263-5_TabPin3" H 7550 1300 50  0001 L CIN
F 3 "http://www.ti.com/lit/ds/symlink/lm2596.pdf" H 7500 1550 50  0001 C CNN
	1    7500 1550
	1    0    0    -1  
$EndComp
Text GLabel 6150 2650 1    50   Input ~ 0
5V
Text GLabel 7000 1150 2    50   Output ~ 0
12V
Text GLabel 6500 1150 0    50   Output ~ 0
GND
Text GLabel 7450 2350 0    50   Input ~ 0
GND
Text GLabel 3450 2300 2    50   Input ~ 0
GND
Text GLabel 3450 2750 2    50   Input ~ 0
GND
Text GLabel 6350 5500 3    50   Input ~ 0
GND
Wire Wire Line
	6800 1450 6800 1650
Text GLabel 2150 3800 2    50   Input ~ 0
12V
Text GLabel 3450 4100 2    50   Input ~ 0
GND
Text GLabel 2150 2900 2    50   Input ~ 0
12V
Text GLabel 2150 3350 2    50   Input ~ 0
12V
Text Label 9050 1650 3    50   ~ 0
5V
$Comp
L Connector_Generic:Conn_01x02 J1
U 1 1 5E6126E2
P 1950 3700
F 0 "J1" V 1914 3512 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 3512 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 3700 50  0001 C CNN
F 3 "~" H 1950 3700 50  0001 C CNN
	1    1950 3700
	-1   0    0    -1  
$EndComp
Text GLabel 3450 3200 2    50   Input ~ 0
GND
Wire Wire Line
	3750 3400 4250 3400
$Comp
L Transistor_FET:IRLZ34N Q6
U 1 1 5E660308
P 3550 3850
F 0 "Q6" H 3754 3896 50  0000 L CNN
F 1 "IRLZ34N" H 3754 3805 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 3775 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 3850 50  0001 L CNN
	1    3550 3850
	-1   0    0    1   
$EndComp
Wire Wire Line
	1100 4050 2150 4050
Wire Wire Line
	1450 4150 1450 4500
Wire Wire Line
	1450 4500 2150 4500
Text Label 1350 4400 3    50   ~ 0
GND_B1
Text Label 1450 4300 3    50   ~ 0
GND_G1
Text Label 1400 4050 2    50   ~ 0
GND_R1
Wire Wire Line
	2150 3600 3450 3600
Wire Wire Line
	2150 3150 3450 3150
Wire Wire Line
	2150 2700 3450 2700
Text Label 4250 3850 2    50   ~ 0
Data_R1
Text Label 4250 4300 2    50   ~ 0
Data_G1
Text Label 4250 4750 2    50   ~ 0
Data_B1
Text Label 4250 2950 2    50   ~ 0
Data_G2
Text Label 4250 2500 2    50   ~ 0
Data_R2
Text Label 4250 3400 2    50   ~ 0
Data_B2
Text Label 3450 3600 0    50   ~ 0
GND_B2
Text Label 3450 3150 0    50   ~ 0
GND_G2
Text Label 3450 2700 0    50   ~ 0
GND_R2
Text Label 8250 1650 2    50   ~ 0
5V_raw
$Comp
L Connector_Generic:Conn_01x02 J7
U 1 1 5E7C68E8
P 9450 2850
F 0 "J7" V 9414 2662 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 9323 2662 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 9450 2850 50  0001 C CNN
F 3 "~" H 9450 2850 50  0001 C CNN
	1    9450 2850
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J6
U 1 1 5E7C76E0
P 9450 3100
F 0 "J6" V 9414 2912 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 9323 2912 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 9450 3100 50  0001 C CNN
F 3 "~" H 9450 3100 50  0001 C CNN
	1    9450 3100
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J8
U 1 1 5E7E2973
P 9450 3600
F 0 "J8" V 9414 3412 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 9323 3412 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 9450 3600 50  0001 C CNN
F 3 "~" H 9450 3600 50  0001 C CNN
	1    9450 3600
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J9
U 1 1 5E7E2D89
P 1950 2450
F 0 "J9" V 1914 2262 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 2262 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 2450 50  0001 C CNN
F 3 "~" H 1950 2450 50  0001 C CNN
	1    1950 2450
	-1   0    0    1   
$EndComp
Text GLabel 9250 3700 0    50   Input ~ 0
GND
Text GLabel 9250 2950 0    50   Input ~ 0
GND
Text GLabel 9250 3200 0    50   Input ~ 0
GND
$Comp
L Connector:Raspberry_Pi_2_3 J11
U 1 1 5E7D03DD
P 6350 4100
F 0 "J11" H 6350 5581 50  0000 C CNN
F 1 "Raspberry_Pi_2_3" H 6350 5490 50  0000 C CNN
F 2 "Module:Raspberry_Pi_Zero_Socketed_THT_FaceDown_MountingHoles" H 6350 4100 50  0001 C CNN
F 3 "https://www.raspberrypi.org/documentation/hardware/raspberrypi/schematics/rpi_SCH_3bplus_1p0_reduced.pdf" H 6350 4100 50  0001 C CNN
	1    6350 4100
	1    0    0    -1  
$EndComp
Text GLabel 9250 3450 0    50   Input ~ 0
GND
$Comp
L Connector_Generic:Conn_01x02 J5
U 1 1 5E7C7DFD
P 9450 3350
F 0 "J5" V 9414 3162 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 9323 3162 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 9450 3350 50  0001 C CNN
F 3 "~" H 9450 3350 50  0001 C CNN
	1    9450 3350
	1    0    0    -1  
$EndComp
NoConn ~ 7150 3300
NoConn ~ 7150 3200
Text GLabel 7150 3800 2    50   Input ~ 0
pin7
Text GLabel 4250 3400 2    50   Input ~ 0
pin7
Text GLabel 5550 3600 0    50   Input ~ 0
pin11
Text GLabel 5550 4400 0    50   Input ~ 0
pin16
Text GLabel 5550 4500 0    50   Input ~ 0
pin18
Text GLabel 5550 4600 0    50   Input ~ 0
pin22
Text GLabel 5550 4800 0    50   Input ~ 0
pin13
Text GLabel 5550 4300 0    50   Input ~ 0
pin15
Text GLabel 7150 4500 2    50   Input ~ 0
pin19
Text GLabel 7150 4400 2    50   Input ~ 0
pin21
Text GLabel 7150 4300 2    50   Input ~ 0
pin24
Text GLabel 5550 4000 0    50   Input ~ 0
pin38
Text GLabel 4250 2950 2    50   Input ~ 0
pin19
Text GLabel 4250 2500 2    50   Input ~ 0
pin21
Text GLabel 4250 3850 2    50   Input ~ 0
pin15
Text GLabel 4250 4300 2    50   Input ~ 0
pin13
Text GLabel 4250 4750 2    50   Input ~ 0
pin11
Text GLabel 8750 3350 0    50   Input ~ 0
pin16
Text Label 8750 3350 0    50   ~ 0
btn16
Text GLabel 8750 3600 0    50   Input ~ 0
pin24
Text GLabel 8750 3100 0    50   Input ~ 0
pin18
Text GLabel 8750 2850 0    50   Input ~ 0
pin22
Text Label 8750 2850 0    50   ~ 0
btn22
Text Label 8750 3100 0    50   ~ 0
btn18
Text Label 3950 2050 0    50   ~ 0
Data_38
Text Label 8750 3600 0    50   ~ 0
btn24
Wire Wire Line
	9250 3600 8750 3600
Wire Wire Line
	9250 3350 8750 3350
Wire Wire Line
	9250 3100 8750 3100
Wire Wire Line
	8750 2850 9250 2850
Wire Wire Line
	3750 4750 4250 4750
Wire Wire Line
	4250 4300 3750 4300
Wire Wire Line
	4250 3850 3750 3850
Wire Wire Line
	4250 2950 3750 2950
Wire Wire Line
	4250 2500 3750 2500
NoConn ~ 5550 3200
NoConn ~ 5550 3300
NoConn ~ 5550 3500
NoConn ~ 5550 3700
NoConn ~ 5550 3900
NoConn ~ 7150 4200
NoConn ~ 7150 3500
NoConn ~ 7150 4600
NoConn ~ 7150 4800
NoConn ~ 7150 4900
NoConn ~ 6450 2800
NoConn ~ 6550 2800
NoConn ~ 5550 4700
NoConn ~ 7150 4000
NoConn ~ 7150 3900
Text GLabel 9650 1650 2    50   Output ~ 0
5V
$Comp
L power:PWR_FLAG #FLG0103
U 1 1 5E92244E
P 9050 1450
F 0 "#FLG0103" H 9050 1525 50  0001 C CNN
F 1 "PWR_FLAG" H 9050 1623 50  0000 C CNN
F 2 "" H 9050 1450 50  0001 C CNN
F 3 "~" H 9050 1450 50  0001 C CNN
	1    9050 1450
	1    0    0    -1  
$EndComp
Connection ~ 9050 1450
Wire Wire Line
	6250 2800 6150 2800
Wire Wire Line
	6150 2800 6150 2650
Connection ~ 6150 2800
Wire Wire Line
	6650 5400 6650 5500
Wire Wire Line
	6650 5500 6350 5500
Wire Wire Line
	6350 5500 6350 5400
NoConn ~ 6450 5400
NoConn ~ 6550 5400
$Comp
L Switch:SW_Push SW1
U 1 1 5E7FD75B
P 7350 3600
F 0 "SW1" H 7350 3885 50  0000 C CNN
F 1 "SW_Push" H 7350 3794 50  0000 C CNN
F 2 "Button_Switch_THT:SW_PUSH_6mm" H 7350 3800 50  0001 C CNN
F 3 "~" H 7350 3800 50  0001 C CNN
	1    7350 3600
	1    0    0    -1  
$EndComp
Text GLabel 7550 3600 2    50   Input ~ 0
GND
NoConn ~ 5950 5400
NoConn ~ 6150 5400
NoConn ~ 6250 5400
Wire Wire Line
	6350 5500 6050 5500
Wire Wire Line
	6050 5500 6050 5400
Connection ~ 6350 5500
Text GLabel 3450 3650 2    50   Input ~ 0
GND
Text GLabel 3450 4550 2    50   Input ~ 0
GND
$Comp
L Connector_Generic:Conn_01x04 J2
U 1 1 5E66A677
P 900 4150
F 0 "J2" V 864 3862 50  0000 R CNN
F 1 "Conn_01x04" V 773 3862 50  0000 R CNN
F 2 "smartlights:PinHeader_1x04_P2.54mm_Vertical_Lock" H 900 4150 50  0001 C CNN
F 3 "~" H 900 4150 50  0001 C CNN
	1    900  4150
	-1   0    0    -1  
$EndComp
Text GLabel 1100 4350 2    50   Input ~ 0
12V
$Comp
L Connector_Generic:Conn_01x02 J13
U 1 1 5E96F43B
P 1950 4600
F 0 "J13" V 1914 4412 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 4412 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 4600 50  0001 C CNN
F 3 "~" H 1950 4600 50  0001 C CNN
	1    1950 4600
	-1   0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J12
U 1 1 5E9A3FCE
P 1950 4150
F 0 "J12" V 1914 3962 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 3962 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 4150 50  0001 C CNN
F 3 "~" H 1950 4150 50  0001 C CNN
	1    1950 4150
	-1   0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J14
U 1 1 5E9A4740
P 1950 5050
F 0 "J14" V 1914 4862 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 4862 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 5050 50  0001 C CNN
F 3 "~" H 1950 5050 50  0001 C CNN
	1    1950 5050
	-1   0    0    -1  
$EndComp
Text GLabel 2150 5150 2    50   Input ~ 0
12V
Text GLabel 2150 4700 2    50   Input ~ 0
12V
Text GLabel 2150 4250 2    50   Input ~ 0
12V
Wire Wire Line
	2150 4150 2150 4050
Connection ~ 2150 4050
Wire Wire Line
	2150 4050 3450 4050
Wire Wire Line
	2150 4600 2150 4500
Connection ~ 2150 4500
Wire Wire Line
	2150 4500 3450 4500
Wire Wire Line
	2150 5050 2150 4950
Connection ~ 2150 4950
Wire Wire Line
	2150 4950 3450 4950
$Comp
L Connector_Generic:Conn_01x04 J15
U 1 1 5EAE26A2
P 900 2800
F 0 "J15" V 864 2512 50  0000 R CNN
F 1 "Conn_01x04" V 773 2512 50  0000 R CNN
F 2 "smartlights:PinHeader_1x04_P2.54mm_Vertical_Lock" H 900 2800 50  0001 C CNN
F 3 "~" H 900 2800 50  0001 C CNN
	1    900  2800
	-1   0    0    -1  
$EndComp
Wire Wire Line
	2150 3700 2150 3600
Wire Wire Line
	2150 3250 2150 3150
Wire Wire Line
	2150 2800 2150 2700
Wire Wire Line
	2150 2700 1100 2700
Connection ~ 2150 2700
Text GLabel 1100 3000 2    50   Input ~ 0
12V
Wire Wire Line
	1100 2900 1350 2900
Wire Wire Line
	1350 2900 1350 3600
Wire Wire Line
	1350 3600 2150 3600
Connection ~ 2150 3600
Wire Wire Line
	2150 3150 1450 3150
Wire Wire Line
	1450 3150 1450 2800
Wire Wire Line
	1450 2800 1100 2800
Connection ~ 2150 3150
Text GLabel 4250 2050 2    50   Input ~ 0
pin38
Text GLabel 5550 4100 0    50   Input ~ 0
pin40
$Comp
L Transistor_FET:IRLZ34N Q8
U 1 1 5EB42F5E
P 3550 2050
F 0 "Q8" H 3754 2096 50  0000 L CNN
F 1 "IRLZ34N" H 3754 2005 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 1975 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 2050 50  0001 L CNN
	1    3550 2050
	-1   0    0    1   
$EndComp
Wire Wire Line
	4250 2050 3750 2050
Wire Wire Line
	3450 2250 2150 2250
Wire Wire Line
	2150 2250 2150 2350
Text GLabel 2150 2450 2    50   Input ~ 0
12V
Text GLabel 3450 1850 2    50   Input ~ 0
GND
$Comp
L Transistor_FET:IRLZ34N Q7
U 1 1 5EB514E4
P 3550 1600
F 0 "Q7" H 3754 1646 50  0000 L CNN
F 1 "IRLZ34N" H 3754 1555 50  0000 L CNN
F 2 "smartlights:TO-252AA" H 3800 1525 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 3550 1600 50  0001 L CNN
	1    3550 1600
	-1   0    0    1   
$EndComp
Text GLabel 3450 1400 2    50   Input ~ 0
GND
Text GLabel 4250 1600 2    50   Input ~ 0
pin40
Text Label 3950 1600 0    50   ~ 0
Data_40
Wire Wire Line
	4250 1600 3750 1600
Wire Wire Line
	3450 1800 2150 1800
Wire Wire Line
	2150 1800 2150 1900
$Comp
L Connector_Generic:Conn_01x02 J16
U 1 1 5EB5855F
P 1950 2000
F 0 "J16" V 1914 1812 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1823 1812 50  0000 R CNN
F 2 "smartlights:PinHeader_1x02_P2.54mm_Vertical_Lock" H 1950 2000 50  0001 C CNN
F 3 "~" H 1950 2000 50  0001 C CNN
	1    1950 2000
	-1   0    0    1   
$EndComp
Text GLabel 2150 2000 2    50   Input ~ 0
12V
Text Label 3450 2250 0    50   ~ 0
GND_38
Text Label 3450 1800 0    50   ~ 0
GND_40
$EndSCHEMATC
