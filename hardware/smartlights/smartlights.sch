EESchema Schematic File Version 4
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title "Smartlights Pi HAT"
Date ""
Rev ""
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Connector_Generic_MountingPin:Conn_02x20_Odd_Even_MountingPin J1
U 1 1 5E600282
P 4550 3500
F 0 "J1" H 4600 4617 50  0000 C CNN
F 1 "Conn_02x20_Odd_Even_MountingPin" H 4600 4526 50  0000 C CNN
F 2 "Connector_PinSocket_2.54mm:PinSocket_2x20_P2.54mm_Vertical" H 4550 3500 50  0001 C CNN
F 3 "~" H 4550 3500 50  0001 C CNN
	1    4550 3500
	1    0    0    -1  
$EndComp
$Comp
L Regulator_Switching:LM2596T-5 U1
U 1 1 5E60D4BA
P 7350 3200
F 0 "U1" H 7350 3567 50  0000 C CNN
F 1 "LM2596T-5" H 7350 3476 50  0000 C CNN
F 2 "Package_TO_SOT_THT:TO-220-5_P3.4x3.7mm_StaggerOdd_Lead3.8mm_Vertical" H 7400 2950 50  0001 L CIN
F 3 "http://www.ti.com/lit/ds/symlink/lm2596.pdf" H 7350 3200 50  0001 C CNN
	1    7350 3200
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 5E60F67B
P 6350 3450
F 0 "C1" H 6465 3496 50  0000 L CNN
F 1 "680µF 12V" H 6465 3405 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D10.0mm_P5.00mm" H 6388 3300 50  0001 C CNN
F 3 "~" H 6350 3450 50  0001 C CNN
	1    6350 3450
	1    0    0    -1  
$EndComp
$Comp
L Device:C C2
U 1 1 5E6170BF
P 8900 3500
F 0 "C2" H 9015 3546 50  0000 L CNN
F 1 "220µF 5V" H 9015 3455 50  0000 L CNN
F 2 "Capacitor_THT:CP_Radial_D6.3mm_P2.50mm" H 8938 3350 50  0001 C CNN
F 3 "~" H 8900 3500 50  0001 C CNN
	1    8900 3500
	1    0    0    -1  
$EndComp
$Comp
L Device:L L1
U 1 1 5E61909A
P 8500 3300
F 0 "L1" V 8319 3300 50  0000 C CNN
F 1 " 33µH" V 8410 3300 50  0000 C CNN
F 2 "Inductor_THT:L_Radial_D12.5mm_P7.00mm_Fastron_09HCP" H 8500 3300 50  0001 C CNN
F 3 "~" H 8500 3300 50  0001 C CNN
	1    8500 3300
	0    1    1    0   
$EndComp
$Comp
L power:GND #PWR02
U 1 1 5E61B63C
P 7300 4050
F 0 "#PWR02" H 7300 3800 50  0001 C CNN
F 1 "GND" H 7305 3877 50  0000 C CNN
F 2 "" H 7300 4050 50  0001 C CNN
F 3 "" H 7300 4050 50  0001 C CNN
	1    7300 4050
	1    0    0    -1  
$EndComp
Wire Wire Line
	6350 2600 6350 2800
Wire Wire Line
	6850 3100 6350 3100
Connection ~ 6350 3100
Wire Wire Line
	6350 3100 6350 3200
Wire Wire Line
	7350 3500 7350 3750
Wire Wire Line
	7350 3750 7300 3750
Wire Wire Line
	7300 3750 7300 4000
Wire Wire Line
	6850 3300 6850 3750
Wire Wire Line
	6850 3750 7300 3750
Connection ~ 7300 3750
Wire Wire Line
	6350 3600 6350 3750
Wire Wire Line
	6350 3750 6850 3750
Connection ~ 6850 3750
Wire Wire Line
	7350 3750 8100 3750
Wire Wire Line
	8100 3750 8100 3650
Connection ~ 7350 3750
Wire Wire Line
	8100 3750 8900 3750
Wire Wire Line
	8900 3750 8900 3650
Connection ~ 8100 3750
Wire Wire Line
	8650 3300 8900 3300
Wire Wire Line
	8900 3300 8900 3350
Wire Wire Line
	8350 3300 8100 3300
Wire Wire Line
	8100 3350 8100 3300
Connection ~ 8100 3300
Wire Wire Line
	8100 3300 7850 3300
Wire Wire Line
	7850 3100 8900 3100
Wire Wire Line
	8900 3100 8900 3300
Connection ~ 8900 3300
Wire Wire Line
	8900 3100 8900 2350
Wire Wire Line
	4850 2350 4850 2600
Connection ~ 8900 3100
$Comp
L Diode:1N5820 D1
U 1 1 5E625494
P 8100 3500
F 0 "D1" V 8054 3579 50  0000 L CNN
F 1 "1N5820" V 8145 3579 50  0000 L CNN
F 2 "Diode_THT:D_DO-201AD_P15.24mm_Horizontal" H 8100 3325 50  0001 C CNN
F 3 "http://www.vishay.com/docs/88526/1n5820.pdf" H 8100 3500 50  0001 C CNN
	1    8100 3500
	0    1    1    0   
$EndComp
Wire Wire Line
	4850 2350 8900 2350
$Comp
L Transistor_FET:IRLZ34N Q1
U 1 1 5E60A5DF
P 5800 3000
F 0 "Q1" H 6004 3046 50  0000 L CNN
F 1 "IRLZ34N" H 6004 2955 50  0000 L CNN
F 2 "Package_TO_SOT_THT:TO-220-3_Vertical" H 6050 2925 50  0001 L CIN
F 3 "http://www.infineon.com/dgdl/irlz34npbf.pdf?fileId=5546d462533600a40153567206892720" H 5800 3000 50  0001 L CNN
	1    5800 3000
	1    0    0    -1  
$EndComp
Connection ~ 6350 3750
Wire Wire Line
	5900 3200 6350 3200
Connection ~ 6350 3200
Wire Wire Line
	6350 3200 6350 3300
Wire Wire Line
	5400 3750 6050 3750
Wire Wire Line
	5600 3000 5150 3000
Wire Wire Line
	5150 3000 5150 2850
Wire Wire Line
	5150 2850 4350 2850
Wire Wire Line
	4350 3800 4350 3750
Wire Wire Line
	4350 3750 5400 3750
Connection ~ 5400 3750
$Comp
L power:+12V #PWR01
U 1 1 5E65062B
P 6350 2600
F 0 "#PWR01" H 6350 2450 50  0001 C CNN
F 1 "+12V" H 6365 2773 50  0000 C CNN
F 2 "" H 6350 2600 50  0001 C CNN
F 3 "" H 6350 2600 50  0001 C CNN
	1    6350 2600
	1    0    0    -1  
$EndComp
NoConn ~ 4850 2700
NoConn ~ 4350 3900
NoConn ~ 4850 3900
NoConn ~ 4350 3400
NoConn ~ 4350 2600
NoConn ~ 4350 2700
NoConn ~ 4350 2800
Wire Wire Line
	4350 2850 4350 2900
NoConn ~ 4850 2900
NoConn ~ 4850 3000
$Comp
L power:PWR_FLAG #FLG0101
U 1 1 5E6578F4
P 6850 2800
F 0 "#FLG0101" H 6850 2875 50  0001 C CNN
F 1 "PWR_FLAG" H 6850 2973 50  0000 C CNN
F 2 "" H 6850 2800 50  0001 C CNN
F 3 "~" H 6850 2800 50  0001 C CNN
	1    6850 2800
	1    0    0    -1  
$EndComp
$Comp
L power:PWR_FLAG #FLG0102
U 1 1 5E658922
P 7500 4000
F 0 "#FLG0102" H 7500 4075 50  0001 C CNN
F 1 "PWR_FLAG" H 7500 4173 50  0000 C CNN
F 2 "" H 7500 4000 50  0001 C CNN
F 3 "~" H 7500 4000 50  0001 C CNN
	1    7500 4000
	1    0    0    -1  
$EndComp
Wire Wire Line
	7500 4000 7300 4000
Connection ~ 7300 4000
Wire Wire Line
	7300 4000 7300 4050
Wire Wire Line
	6850 2800 6350 2800
Connection ~ 6350 2800
Wire Wire Line
	6350 2800 6350 3100
NoConn ~ 4850 2800
NoConn ~ 4850 3100
NoConn ~ 4850 3200
NoConn ~ 4850 3300
NoConn ~ 4850 3400
NoConn ~ 4850 3500
NoConn ~ 4850 3700
NoConn ~ 4850 3800
NoConn ~ 4850 3600
NoConn ~ 4850 4000
NoConn ~ 4850 4100
NoConn ~ 4850 4200
NoConn ~ 4850 4300
NoConn ~ 4850 4400
NoConn ~ 4850 4500
NoConn ~ 4350 4500
NoConn ~ 4350 4400
NoConn ~ 4350 4300
NoConn ~ 4350 4200
NoConn ~ 4350 4100
NoConn ~ 4350 4000
NoConn ~ 4350 3700
NoConn ~ 4350 3600
NoConn ~ 4350 3500
NoConn ~ 4350 3200
NoConn ~ 4350 3100
NoConn ~ 4350 3000
NoConn ~ 4350 3300
NoConn ~ 4600 4700
$Comp
L Connector:Conn_Coaxial_Power J3
U 1 1 5E60BF44
P 6250 2800
F 0 "J3" V 6033 2750 50  0000 C CNN
F 1 "Conn_Coaxial_Power" V 6124 2750 50  0000 C CNN
F 2 "Connector_BarrelJack:BarrelJack_Horizontal" H 6250 2750 50  0001 C CNN
F 3 "~" H 6250 2750 50  0001 C CNN
	1    6250 2800
	0    1    1    0   
$EndComp
Wire Wire Line
	6050 2800 6050 3750
Connection ~ 6050 3750
Wire Wire Line
	6050 3750 6350 3750
$Comp
L Connector_Generic_MountingPin:Conn_01x02_MountingPin J2
U 1 1 5E6126E2
P 5550 2500
F 0 "J2" V 5775 2368 50  0000 C CNN
F 1 "Conn_01x02_MountingPin" V 5684 2368 50  0000 C CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x02_P2.54mm_Vertical" H 5550 2500 50  0001 C CNN
F 3 "~" H 5550 2500 50  0001 C CNN
	1    5550 2500
	0    -1   -1   0   
$EndComp
Wire Wire Line
	5550 2700 5400 2700
Wire Wire Line
	5400 2700 5400 3750
Wire Wire Line
	5650 2700 5900 2700
Wire Wire Line
	5900 2700 5900 2800
NoConn ~ 5850 2500
$EndSCHEMATC
