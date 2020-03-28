EESchema Schematic File Version 4
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "Smartlights Pi HAT"
Date ""
Rev "2.1"
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Transistor_FET:IRLZ34N Q4
U 1 1 5E6648A2
P 1800 4300
F 0 "Q4" H 2004 4346 50  0000 L CNN
F 1 "IRLZ34N" H 2004 4255 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 4225 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 4300 50  0001 L CNN
	1    1800 4300
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q2
U 1 1 5E66544B
P 1800 4750
F 0 "Q2" H 2004 4796 50  0000 L CNN
F 1 "IRLZ34N" H 2004 4705 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 4675 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 4750 50  0001 L CNN
	1    1800 4750
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q3
U 1 1 5E665F41
P 1800 2650
F 0 "Q3" H 2004 2696 50  0000 L CNN
F 1 "IRLZ34N" H 2004 2605 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 2575 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 2650 50  0001 L CNN
	1    1800 2650
	-1   0    0    1   
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q5
U 1 1 5E66772F
P 1800 3250
F 0 "Q5" H 2004 3296 50  0000 L CNN
F 1 "IRLZ34N" H 2004 3205 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 3175 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 3250 50  0001 L CNN
	1    1800 3250
	-1   0    0    1   
$EndComp
$Comp
L Connector_Generic:Conn_01x04 J2
U 1 1 5E66A677
P 900 4250
F 0 "J2" V 864 3962 50  0000 R CNN
F 1 "Conn_01x04" V 773 3962 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x04_P2.54mm_Vertical" H 900 4250 50  0001 C CNN
F 3 "~" H 900 4250 50  0001 C CNN
	1    900  4250
	-1   0    0    1   
$EndComp
Wire Wire Line
	1100 4250 1350 4250
Wire Wire Line
	1350 4250 1350 4950
Wire Wire Line
	1350 4950 1700 4950
Wire Wire Line
	1450 4150 1100 4150
$Comp
L Connector_Generic:Conn_01x02 J3
U 1 1 5E67B518
P 1400 3550
F 0 "J3" V 1364 3362 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1273 3362 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 1400 3550 50  0001 C CNN
F 3 "~" H 1400 3550 50  0001 C CNN
	1    1400 3550
	-1   0    0    1   
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J4
U 1 1 5E67BE71
P 1400 2950
F 0 "J4" V 1364 2762 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1273 2762 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 1400 2950 50  0001 C CNN
F 3 "~" H 1400 2950 50  0001 C CNN
	1    1400 2950
	-1   0    0    1   
$EndComp
Wire Wire Line
	7300 1650 7900 1650
$Comp
L Connector:Conn_Coaxial_Power J10
U 1 1 5E60BF44
P 4950 1150
F 0 "J10" V 4733 1100 50  0000 C CNN
F 1 "Conn_Coaxial_Power" V 4824 1100 50  0000 C CNN
F 2 "Connector_BarrelJack:BarrelJack_Horizontal" H 4950 1100 50  0001 C CNN
F 3 "~" H 4950 1100 50  0001 C CNN
	1    4950 1150
	0    1    1    0   
$EndComp
Wire Wire Line
	5050 1150 5050 1450
Connection ~ 5050 1150
Wire Wire Line
	5250 1150 5050 1150
Wire Wire Line
	5700 2350 5700 2400
Connection ~ 5700 2350
Wire Wire Line
	5900 2350 5700 2350
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5E658922
P 5900 2350
F 0 "#FLG0102" H 5900 2425 50  0001 C CNN
F 1 "PWR_FLAG" H 5900 2523 50  0000 C CNN
F 2 "" H 5900 2350 50  0001 C CNN
F 3 "~" H 5900 2350 50  0001 C CNN
	1    5900 2350
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5E6578F4
P 5050 1450
F 0 "#FLG0101" H 5050 1525 50  0001 C CNN
F 1 "PWR_FLAG" H 5050 1623 50  0000 C CNN
F 2 "" H 5050 1450 50  0001 C CNN
F 3 "~" H 5050 1450 50  0001 C CNN
	1    5050 1450
	0    -1   -1   0   
$EndComp
$Comp
L power:+12V #PWR01
U 1 1 5E65062B
P 5050 950
F 0 "#PWR01" H 5050 800 50  0001 C CNN
F 1 "+12V" H 5065 1123 50  0000 C CNN
F 2 "" H 5050 950 50  0001 C CNN
F 3 "" H 5050 950 50  0001 C CNN
	1    5050 950 
	1    0    0    -1  
$EndComp
$Comp
L Transistor_FET:IRLZ34N Q1
U 1 1 5E60A5DF
P 1800 5600
F 0 "Q1" H 2004 5646 50  0000 L CNN
F 1 "IRLZ34N" H 2004 5555 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 5525 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 5600 50  0001 L CNN
	1    1800 5600
	-1   0    0    1   
$EndComp
$Comp
L Diode:1N5820 D1
U 1 1 5E625494
P 6500 1850
F 0 "D1" V 6454 1929 50  0000 L CNN
F 1 "1N5820" V 6545 1929 50  0000 L CNN
F 2 "Diode_THT:D_DO-201AD_P15.24mm_Horizontal" H 6500 1675 50  0001 C CNN
F 3 "http://www.vishay.com/docs/88526/1n5820.pdf" H 6500 1850 50  0001 C CNN
	1    6500 1850
	0    1    1    0   
$EndComp
Connection ~ 7300 1650
Wire Wire Line
	7300 1450 7300 1650
Wire Wire Line
	6250 1450 7300 1450
Wire Wire Line
	6500 1650 6250 1650
Connection ~ 6500 1650
Wire Wire Line
	6500 1700 6500 1650
Wire Wire Line
	6750 1650 6500 1650
Wire Wire Line
	7300 1650 7300 1700
Wire Wire Line
	7050 1650 7300 1650
Connection ~ 6500 2100
Wire Wire Line
	7300 2100 7300 2000
Wire Wire Line
	6500 2100 7300 2100
Connection ~ 5750 2100
Wire Wire Line
	6500 2100 6500 2000
Wire Wire Line
	5750 2100 6500 2100
Connection ~ 5250 2100
Wire Wire Line
	5050 2100 5250 2100
Wire Wire Line
	5050 1950 5050 2100
Connection ~ 5700 2100
Wire Wire Line
	5250 2100 5700 2100
Wire Wire Line
	5250 1650 5250 2100
Wire Wire Line
	5700 2100 5700 2350
Wire Wire Line
	5750 2100 5700 2100
Wire Wire Line
	5750 1850 5750 2100
Connection ~ 5050 1450
Wire Wire Line
	5250 1450 5050 1450
Wire Wire Line
	5050 950  5050 1150
$Comp
L power:GND #PWR02
U 1 1 5E61B63C
P 5700 2400
F 0 "#PWR02" H 5700 2150 50  0001 C CNN
F 1 "GND" H 5705 2227 50  0000 C CNN
F 2 "" H 5700 2400 50  0001 C CNN
F 3 "" H 5700 2400 50  0001 C CNN
	1    5700 2400
	1    0    0    -1  
$EndComp
$Comp
L Device:L L1
U 1 1 5E61909A
P 6900 1650
F 0 "L1" V 6719 1650 50  0000 C CNN
F 1 " 33µH" V 6810 1650 50  0000 C CNN
F 2 "Inductor_THT:L_Radial_D12.5mm_P7.00mm_Fastron_09HCP" H 6900 1650 50  0001 C CNN
F 3 "~" H 6900 1650 50  0001 C CNN
	1    6900 1650
	0    1    1    0   
$EndComp
$Comp
L Device:C C2
U 1 1 5E6170BF
P 7300 1850
F 0 "C2" H 7415 1896 50  0000 L CNN
F 1 "220µF 5V" H 7415 1805 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D6.3mm_P2.50mm" H 7338 1700 50  0001 C CNN
F 3 "~" H 7300 1850 50  0001 C CNN
	1    7300 1850
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 5E60F67B
P 5050 1800
F 0 "C1" H 5165 1846 50  0000 L CNN
F 1 "680µF 12V" H 5165 1755 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D10.0mm_P5.00mm" H 5088 1650 50  0001 C CNN
F 3 "~" H 5050 1800 50  0001 C CNN
	1    5050 1800
	1    0    0    -1  
$EndComp
$Comp
L Regulator_Switching:LM2596T-5 U1
U 1 1 5E60D4BA
P 5750 1550
F 0 "U1" H 5750 1917 50  0000 C CNN
F 1 "LM2596T-5" H 5750 1826 50  0000 C CNN
F 2 "Package_TO_SOT_THT:TO-220-5_P3.4x3.7mm_StaggerOdd_Lead3.8mm_Vertical" H 5800 1300 50  0001 L CIN
F 3 "http://www.ti.com/lit/ds/symlink/lm2596.pdf" H 5750 1550 50  0001 C CNN
	1    5750 1550
	1    0    0    -1  
$EndComp
Text GLabel 4400 2650 1    50   Input ~ 0
5V
Text GLabel 5250 1150 2    50   Output ~ 0
12V
Text GLabel 4750 1150 0    50   Output ~ 0
GND
Text GLabel 5700 2350 0    50   Input ~ 0
GND
Text GLabel 1700 2950 2    50   Input ~ 0
GND
Text GLabel 1100 4350 2    50   Input ~ 0
GND
Text GLabel 1700 3550 2    50   Input ~ 0
GND
Text GLabel 4600 5500 3    50   Input ~ 0
GND
Wire Wire Line
	5050 1450 5050 1650
Text GLabel 1700 5400 2    50   Input ~ 0
12V
Text GLabel 1700 3650 2    50   Input ~ 0
12V
Text GLabel 1700 4100 2    50   Input ~ 0
12V
Text GLabel 1700 4550 2    50   Input ~ 0
12V
Text GLabel 1700 2450 2    50   Input ~ 0
12V
Text GLabel 1700 3050 2    50   Input ~ 0
12V
Text Label 7300 1650 3    50   ~ 0
5V
$Comp
L Connector_Generic:Conn_01x02 J1
U 1 1 5E6126E2
P 1400 5900
F 0 "J1" V 1364 5712 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 1273 5712 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 1400 5900 50  0001 C CNN
F 3 "~" H 1400 5900 50  0001 C CNN
	1    1400 5900
	-1   0    0    1   
$EndComp
Text GLabel 1700 5900 2    50   Input ~ 0
GND
Wire Wire Line
	2000 5600 2500 5600
$Comp
L Transistor_FET:IRLZ34N Q6
U 1 1 5E660308
P 1800 3850
F 0 "Q6" H 2004 3896 50  0000 L CNN
F 1 "IRLZ34N" H 2004 3805 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 2050 3775 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 1800 3850 50  0001 L CNN
	1    1800 3850
	-1   0    0    1   
$EndComp
Wire Wire Line
	1100 4050 1700 4050
Wire Wire Line
	1450 4150 1450 4500
Wire Wire Line
	1450 4500 1700 4500
Text Label 1350 4400 3    50   ~ 0
PWR_R
Text Label 1450 4300 3    50   ~ 0
PWR_G
Text Label 1400 4050 2    50   ~ 0
PWR_B
Wire Wire Line
	1600 5800 1700 5800
Wire Wire Line
	1700 5900 1600 5900
Wire Wire Line
	1700 3550 1600 3550
Wire Wire Line
	1600 3450 1700 3450
Wire Wire Line
	1700 2950 1600 2950
Wire Wire Line
	1600 2850 1700 2850
Text Label 2500 3850 2    50   ~ 0
Data_B
Text Label 2500 4300 2    50   ~ 0
Data_G
Text Label 2500 4750 2    50   ~ 0
Data_R
Text Label 2500 3250 2    50   ~ 0
Data_19
Text Label 2500 2650 2    50   ~ 0
Data_21
Text Label 2500 5600 2    50   ~ 0
Data_7
Text Label 1700 5800 0    50   ~ 0
PWR_7
Text Label 1700 3450 0    50   ~ 0
PWR_19
Text Label 1700 2850 0    50   ~ 0
PWR_21
Text Label 6500 1650 2    50   ~ 0
5V_raw
$Comp
L Connector_Generic:Conn_01x02 J7
U 1 1 5E7C68E8
P 7700 2850
F 0 "J7" V 7664 2662 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 7573 2662 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 7700 2850 50  0001 C CNN
F 3 "~" H 7700 2850 50  0001 C CNN
	1    7700 2850
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J6
U 1 1 5E7C76E0
P 7700 3100
F 0 "J6" V 7664 2912 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 7573 2912 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 7700 3100 50  0001 C CNN
F 3 "~" H 7700 3100 50  0001 C CNN
	1    7700 3100
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J8
U 1 1 5E7E2973
P 7700 3850
F 0 "J8" V 7664 3662 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 7573 3662 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 7700 3850 50  0001 C CNN
F 3 "~" H 7700 3850 50  0001 C CNN
	1    7700 3850
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J9
U 1 1 5E7E2D89
P 7700 3600
F 0 "J9" V 7664 3412 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 7573 3412 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 7700 3600 50  0001 C CNN
F 3 "~" H 7700 3600 50  0001 C CNN
	1    7700 3600
	1    0    0    -1  
$EndComp
Text GLabel 7500 3700 0    50   Input ~ 0
GND
Text GLabel 7500 3950 0    50   Input ~ 0
GND
Text GLabel 7500 2950 0    50   Input ~ 0
GND
Text GLabel 7500 3200 0    50   Input ~ 0
GND
$Comp
L Connector:Raspberry_Pi_2_3 J11
U 1 1 5E7D03DD
P 4600 4100
F 0 "J11" H 4600 5581 50  0000 C CNN
F 1 "Raspberry_Pi_2_3" H 4600 5490 50  0000 C CNN
F 2 "Module:Raspberry_Pi_Zero_Socketed_THT_FaceDown_MountingHoles" H 4600 4100 50  0001 C CNN
F 3 "https://www.raspberrypi.org/documentation/hardware/raspberrypi/schematics/rpi_SCH_3bplus_1p0_reduced.pdf" H 4600 4100 50  0001 C CNN
	1    4600 4100
	1    0    0    -1  
$EndComp
Text GLabel 7500 3450 0    50   Input ~ 0
GND
$Comp
L Connector_Generic:Conn_01x02 J5
U 1 1 5E7C7DFD
P 7700 3350
F 0 "J5" V 7664 3162 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 7573 3162 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 7700 3350 50  0001 C CNN
F 3 "~" H 7700 3350 50  0001 C CNN
	1    7700 3350
	1    0    0    -1  
$EndComp
NoConn ~ 5400 3300
NoConn ~ 5400 3200
Text GLabel 5400 3800 2    50   Input ~ 0
pin7
Text GLabel 2500 5600 2    50   Input ~ 0
pin7
Text GLabel 3800 3600 0    50   Input ~ 0
pin11
Text GLabel 3800 4400 0    50   Input ~ 0
pin16
Text GLabel 3800 4500 0    50   Input ~ 0
pin18
Text GLabel 3800 4600 0    50   Input ~ 0
pin22
Text GLabel 3800 4800 0    50   Input ~ 0
pin13
Text GLabel 3800 4300 0    50   Input ~ 0
pin15
Text GLabel 5400 4500 2    50   Input ~ 0
pin19
Text GLabel 5400 4400 2    50   Input ~ 0
pin21
Text GLabel 5400 4300 2    50   Input ~ 0
pin24
Text GLabel 5400 4200 2    50   Input ~ 0
pin26
Text GLabel 2500 3250 2    50   Input ~ 0
pin19
Text GLabel 2500 2650 2    50   Input ~ 0
pin21
Text GLabel 2500 3850 2    50   Input ~ 0
pin15
Text GLabel 2500 4300 2    50   Input ~ 0
pin13
Text GLabel 2500 4750 2    50   Input ~ 0
pin11
Text GLabel 7000 3350 0    50   Input ~ 0
pin16
Text Label 7000 3350 0    50   ~ 0
btn16
Text GLabel 7000 3850 0    50   Input ~ 0
pin24
Text GLabel 7000 3600 0    50   Input ~ 0
pin26
Text GLabel 7000 3100 0    50   Input ~ 0
pin18
Text GLabel 7000 2850 0    50   Input ~ 0
pin22
Text Label 7000 2850 0    50   ~ 0
btn22
Text Label 7000 3100 0    50   ~ 0
btn18
Text Label 7000 3600 0    50   ~ 0
btn26
Text Label 7000 3850 0    50   ~ 0
btn24
Wire Wire Line
	7500 3600 7000 3600
Wire Wire Line
	7500 3850 7000 3850
Wire Wire Line
	7500 3350 7000 3350
Wire Wire Line
	7500 3100 7000 3100
Wire Wire Line
	7000 2850 7500 2850
Wire Wire Line
	2000 4750 2500 4750
Wire Wire Line
	2500 4300 2000 4300
Wire Wire Line
	2500 3850 2000 3850
Wire Wire Line
	2500 3250 2000 3250
Wire Wire Line
	2500 2650 2000 2650
NoConn ~ 3800 3200
NoConn ~ 3800 3300
NoConn ~ 3800 3500
NoConn ~ 3800 3700
NoConn ~ 3800 3900
NoConn ~ 3800 4000
NoConn ~ 3800 4100
NoConn ~ 5400 3500
NoConn ~ 5400 4600
NoConn ~ 5400 4800
NoConn ~ 5400 4900
NoConn ~ 4700 2800
NoConn ~ 4800 2800
NoConn ~ 3800 4700
NoConn ~ 5400 4000
NoConn ~ 5400 3900
Text GLabel 7900 1650 2    50   Output ~ 0
5V
$Comp
L power:PWR_FLAG #FLG0103
U 1 1 5E92244E
P 7300 1450
F 0 "#FLG0103" H 7300 1525 50  0001 C CNN
F 1 "PWR_FLAG" H 7300 1623 50  0000 C CNN
F 2 "" H 7300 1450 50  0001 C CNN
F 3 "~" H 7300 1450 50  0001 C CNN
	1    7300 1450
	1    0    0    -1  
$EndComp
Connection ~ 7300 1450
Wire Wire Line
	4500 2800 4400 2800
Wire Wire Line
	4400 2800 4400 2650
Connection ~ 4400 2800
Wire Wire Line
	4900 5400 4900 5500
Wire Wire Line
	4900 5500 4600 5500
Wire Wire Line
	4600 5500 4600 5400
NoConn ~ 4700 5400
NoConn ~ 4800 5400
$Comp
L Switch:SW_Push SW1
U 1 1 5E7FD75B
P 5600 3600
F 0 "SW1" H 5600 3885 50  0000 C CNN
F 1 "SW_Push" H 5600 3794 50  0000 C CNN
F 2 "Button_Switch_THT:SW_PUSH_6mm" H 5600 3800 50  0001 C CNN
F 3 "~" H 5600 3800 50  0001 C CNN
	1    5600 3600
	1    0    0    -1  
$EndComp
Text GLabel 5800 3600 2    50   Input ~ 0
GND
NoConn ~ 4200 5400
NoConn ~ 4400 5400
NoConn ~ 4500 5400
Wire Wire Line
	4600 5500 4300 5500
Wire Wire Line
	4300 5500 4300 5400
Connection ~ 4600 5500
$EndSCHEMATC
