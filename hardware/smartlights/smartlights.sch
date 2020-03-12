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
NoConn ~ 3900 2650
NoConn ~ 3900 3850
NoConn ~ 3400 2550
NoConn ~ 3400 2650
NoConn ~ 3400 2750
NoConn ~ 3900 2850
NoConn ~ 3900 2950
NoConn ~ 3900 2750
NoConn ~ 3900 3050
NoConn ~ 3400 2850
NoConn ~ 3900 3750
NoConn ~ 3900 3950
NoConn ~ 3900 4050
NoConn ~ 3900 4250
NoConn ~ 3900 2550
NoConn ~ 3900 2450
NoConn ~ 3400 2450
NoConn ~ 3400 4350
NoConn ~ 3400 4250
NoConn ~ 3400 4150
NoConn ~ 3400 3550
NoConn ~ 3400 3150
NoConn ~ 3400 3050
NoConn ~ 3400 2950
NoConn ~ 3400 3250
NoConn ~ 3900 4150
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
	2000 4750 2900 4750
Wire Wire Line
	2900 3850 3400 3850
Wire Wire Line
	2700 3850 2700 3650
Wire Wire Line
	2700 3650 3400 3650
Wire Wire Line
	7800 3200 8400 3200
$Comp
L Connector:Conn_Coaxial_Power J10
U 1 1 5E60BF44
P 5450 2700
F 0 "J10" V 5233 2650 50  0000 C CNN
F 1 "Conn_Coaxial_Power" V 5324 2650 50  0000 C CNN
F 2 "Connector_BarrelJack:BarrelJack_Horizontal" H 5450 2650 50  0001 C CNN
F 3 "~" H 5450 2650 50  0001 C CNN
	1    5450 2700
	0    1    1    0   
$EndComp
Wire Wire Line
	5550 2700 5550 3000
Connection ~ 5550 2700
Wire Wire Line
	5750 2700 5550 2700
Wire Wire Line
	6200 3900 6200 3950
Connection ~ 6200 3900
Wire Wire Line
	6400 3900 6200 3900
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5E658922
P 6400 3900
F 0 "#FLG0102" H 6400 3975 50  0001 C CNN
F 1 "PWR_FLAG" H 6400 4073 50  0000 C CNN
F 2 "" H 6400 3900 50  0001 C CNN
F 3 "~" H 6400 3900 50  0001 C CNN
	1    6400 3900
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5E6578F4
P 5550 3000
F 0 "#FLG0101" H 5550 3075 50  0001 C CNN
F 1 "PWR_FLAG" H 5550 3173 50  0000 C CNN
F 2 "" H 5550 3000 50  0001 C CNN
F 3 "~" H 5550 3000 50  0001 C CNN
	1    5550 3000
	0    -1   -1   0   
$EndComp
$Comp
L power:+12V #PWR01
U 1 1 5E65062B
P 5550 2500
F 0 "#PWR01" H 5550 2350 50  0001 C CNN
F 1 "+12V" H 5565 2673 50  0000 C CNN
F 2 "" H 5550 2500 50  0001 C CNN
F 3 "" H 5550 2500 50  0001 C CNN
	1    5550 2500
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
P 7000 3400
F 0 "D1" V 6954 3479 50  0000 L CNN
F 1 "1N5820" V 7045 3479 50  0000 L CNN
F 2 "Diode_THT:D_DO-201AD_P15.24mm_Horizontal" H 7000 3225 50  0001 C CNN
F 3 "http://www.vishay.com/docs/88526/1n5820.pdf" H 7000 3400 50  0001 C CNN
	1    7000 3400
	0    1    1    0   
$EndComp
Connection ~ 7800 3200
Wire Wire Line
	7800 3000 7800 3200
Wire Wire Line
	6750 3000 7800 3000
Wire Wire Line
	7000 3200 6750 3200
Connection ~ 7000 3200
Wire Wire Line
	7000 3250 7000 3200
Wire Wire Line
	7250 3200 7000 3200
Wire Wire Line
	7800 3200 7800 3250
Wire Wire Line
	7550 3200 7800 3200
Connection ~ 7000 3650
Wire Wire Line
	7800 3650 7800 3550
Wire Wire Line
	7000 3650 7800 3650
Connection ~ 6250 3650
Wire Wire Line
	7000 3650 7000 3550
Wire Wire Line
	6250 3650 7000 3650
Connection ~ 5750 3650
Wire Wire Line
	5550 3650 5750 3650
Wire Wire Line
	5550 3500 5550 3650
Connection ~ 6200 3650
Wire Wire Line
	5750 3650 6200 3650
Wire Wire Line
	5750 3200 5750 3650
Wire Wire Line
	6200 3650 6200 3900
Wire Wire Line
	6250 3650 6200 3650
Wire Wire Line
	6250 3400 6250 3650
Connection ~ 5550 3000
Wire Wire Line
	5750 3000 5550 3000
Wire Wire Line
	5550 2500 5550 2700
$Comp
L power:GND #PWR02
U 1 1 5E61B63C
P 6200 3950
F 0 "#PWR02" H 6200 3700 50  0001 C CNN
F 1 "GND" H 6205 3777 50  0000 C CNN
F 2 "" H 6200 3950 50  0001 C CNN
F 3 "" H 6200 3950 50  0001 C CNN
	1    6200 3950
	1    0    0    -1  
$EndComp
$Comp
L Device:L L1
U 1 1 5E61909A
P 7400 3200
F 0 "L1" V 7219 3200 50  0000 C CNN
F 1 " 33µH" V 7310 3200 50  0000 C CNN
F 2 "Inductor_THT:L_Radial_D12.5mm_P7.00mm_Fastron_09HCP" H 7400 3200 50  0001 C CNN
F 3 "~" H 7400 3200 50  0001 C CNN
	1    7400 3200
	0    1    1    0   
$EndComp
$Comp
L Device:C C2
U 1 1 5E6170BF
P 7800 3400
F 0 "C2" H 7915 3446 50  0000 L CNN
F 1 "220µF 5V" H 7915 3355 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D6.3mm_P2.50mm" H 7838 3250 50  0001 C CNN
F 3 "~" H 7800 3400 50  0001 C CNN
	1    7800 3400
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 5E60F67B
P 5550 3350
F 0 "C1" H 5665 3396 50  0000 L CNN
F 1 "680µF 12V" H 5665 3305 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D10.0mm_P5.00mm" H 5588 3200 50  0001 C CNN
F 3 "~" H 5550 3350 50  0001 C CNN
	1    5550 3350
	1    0    0    -1  
$EndComp
$Comp
L Regulator_Switching:LM2596T-5 U1
U 1 1 5E60D4BA
P 6250 3100
F 0 "U1" H 6250 3467 50  0000 C CNN
F 1 "LM2596T-5" H 6250 3376 50  0000 C CNN
F 2 "Package_TO_SOT_THT:TO-220-5_P3.4x3.7mm_StaggerOdd_Lead3.8mm_Vertical" H 6300 2850 50  0001 L CIN
F 3 "http://www.ti.com/lit/ds/symlink/lm2596.pdf" H 6250 3100 50  0001 C CNN
	1    6250 3100
	1    0    0    -1  
$EndComp
Text GLabel 8400 3200 2    50   Output ~ 0
5V
Text GLabel 3900 4350 2    50   Input ~ 0
5V
Text GLabel 5750 2700 2    50   Output ~ 0
12V
Text GLabel 5250 2700 0    50   Output ~ 0
GND
Text GLabel 6200 3900 0    50   Input ~ 0
GND
Text GLabel 1700 2950 2    50   Input ~ 0
GND
Text GLabel 1100 4350 2    50   Input ~ 0
GND
Text GLabel 1700 3550 2    50   Input ~ 0
GND
Text GLabel 3400 3950 0    50   Input ~ 0
GND
Wire Wire Line
	5550 3000 5550 3200
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
Wire Wire Line
	3400 3350 3000 3350
Wire Wire Line
	3000 3350 3000 2650
Wire Wire Line
	3000 2650 2000 2650
Wire Wire Line
	3400 3450 2900 3450
Wire Wire Line
	2900 3450 2900 3250
Wire Wire Line
	2900 3250 2000 3250
Text Label 7800 3200 3    50   ~ 0
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
	2000 5600 3000 5600
Wire Wire Line
	3000 5600 3000 4050
Wire Wire Line
	3000 4050 3400 4050
Wire Wire Line
	2000 3850 2700 3850
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
Wire Wire Line
	2900 3850 2900 4750
Wire Wire Line
	3400 3750 2800 3750
Wire Wire Line
	2800 3750 2800 4300
Wire Wire Line
	2800 4300 2000 4300
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
Text Label 3300 3650 2    50   ~ 0
Data_B
Text Label 3300 3750 2    50   ~ 0
Data_G
Text Label 3300 3850 2    50   ~ 0
Data_R
Text Label 3350 3450 2    50   ~ 0
Data_19
Text Label 3350 3350 2    50   ~ 0
Data_21
Text Label 3300 4050 2    50   ~ 0
Data_7
Text Label 1700 5800 0    50   ~ 0
PWR_7
Text Label 1700 3450 0    50   ~ 0
PWR_19
Text Label 1700 2850 0    50   ~ 0
PWR,21
Text Label 7000 3200 2    50   ~ 0
5V_raw
$Comp
L Connector_Generic:Conn_02x20_Odd_Even J11
U 1 1 5E600282
P 3600 3450
F 0 "J11" H 3650 4567 50  0000 C CNN
F 1 "Conn_02x20_Odd_Even_MountingPin" H 3650 4476 50  0000 C CNN
F 2 "Connector_PinSocket_2.54mm:PinSocket_2x20_P2.54mm_Vertical" H 3600 3450 50  0001 C CNN
F 3 "~" H 3600 3450 50  0001 C CNN
	1    3600 3450
	1    0    0    1   
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J7
U 1 1 5E7C68E8
P 4650 3350
F 0 "J7" V 4614 3162 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 4523 3162 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 4650 3350 50  0001 C CNN
F 3 "~" H 4650 3350 50  0001 C CNN
	1    4650 3350
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J6
U 1 1 5E7C76E0
P 4650 3600
F 0 "J6" V 4614 3412 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 4523 3412 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 4650 3600 50  0001 C CNN
F 3 "~" H 4650 3600 50  0001 C CNN
	1    4650 3600
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J5
U 1 1 5E7C7DFD
P 4650 3850
F 0 "J5" V 4614 3662 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 4523 3662 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 4650 3850 50  0001 C CNN
F 3 "~" H 4650 3850 50  0001 C CNN
	1    4650 3850
	1    0    0    -1  
$EndComp
Wire Wire Line
	4450 3350 3900 3350
$Comp
L Connector_Generic:Conn_01x02 J8
U 1 1 5E7E2973
P 4650 3100
F 0 "J8" V 4614 2912 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 4523 2912 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 4650 3100 50  0001 C CNN
F 3 "~" H 4650 3100 50  0001 C CNN
	1    4650 3100
	1    0    0    -1  
$EndComp
$Comp
L Connector_Generic:Conn_01x02 J9
U 1 1 5E7E2D89
P 4650 2850
F 0 "J9" V 4614 2662 50  0000 R CNN
F 1 "Conn_01x02_MountingPin" V 4523 2662 50  0000 R CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 4650 2850 50  0001 C CNN
F 3 "~" H 4650 2850 50  0001 C CNN
	1    4650 2850
	1    0    0    -1  
$EndComp
Text GLabel 4450 2950 0    50   Input ~ 0
GND
NoConn ~ 3900 3450
Text GLabel 4450 3200 0    50   Input ~ 0
GND
Text GLabel 4450 3450 0    50   Input ~ 0
GND
Text GLabel 4450 3700 0    50   Input ~ 0
GND
Text GLabel 4450 3950 0    50   Input ~ 0
GND
Wire Wire Line
	4100 3150 4100 2850
Wire Wire Line
	3900 3150 4100 3150
Wire Wire Line
	4100 2850 4450 2850
Wire Wire Line
	4200 3100 4200 3250
Wire Wire Line
	4450 3100 4200 3100
Wire Wire Line
	4200 3250 3900 3250
Wire Wire Line
	4100 3850 4100 3650
Wire Wire Line
	4450 3850 4100 3850
Wire Wire Line
	4100 3650 3900 3650
Wire Wire Line
	4200 3550 3900 3550
Wire Wire Line
	4450 3600 4200 3600
Wire Wire Line
	4200 3600 4200 3550
$EndSCHEMATC
