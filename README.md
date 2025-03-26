# Fourier Series

A complete Go program that demonstrates a Fourier Series approximation of a square wave.

## Overview

This repository contains a complete Go program that demonstrates a Fourier Series approximation of a square wave. The example includes generating the series, calculating coefficients, and visualizing the result using a simple console output.

## Features

This program:

### Structures
- Defines basic structures:
  - `Point` for x,y coordinates
  - `FourierSeries` for series parameters

### Calculations
- Implements Fourier Series calculations:
  - `CalculateCoefficient`: Computes coefficients for a square wave (only odd harmonics)
  - `Evaluate`: Calculates the series value at a given time
  - `GeneratePoints`: Creates sample points across one period

### Visualization
- Includes a simple visualization:
  - `Visualize`: Creates a basic console-based plot using ASCII characters
  - Shows the waveform with a zero line reference

## Execution

When you run this program, it will:
1. Calculate a Fourier series approximation of a square wave
2. Show some sample values
3. Display a simple ASCII visualization

## Sample Output

Example values:
t=0.00, f(t)=0.0000
t=0.40, f(t)=1.2732
t=0.80, f(t)=0.3183
t=1.20, f(t)=-0.3183
t=1.60, f(t)=-1.2732

Fourier Series Approximation of Square Wave:
```
   *                         *                                       
   ***      ***      **      ***                                      
  *   **  **  *** ***  **  **   *                                     
  *    ***      ***      ***    *                                     
 *                               *                                    
 *                               *                                    
 *                               *                                    
*                                 *                                   
                                                                      
*                                 *                                   
*----------------------------------*----------------------------------
                                   *                                 *
                                                                      
                                   *                                 *
                                    *                               * 
                                    *                               * 
                                    *                               * 
                                     *    ***      ***      ***    *  
                                     *   **  **  *** ***  **  **   *  
                                      ***      ***      **      ***   
```     

Using 7 harmonics
This approximates a square wave using sine terms.


## Customization

This is a basic implementation that approximates a square wave. You can modify:
- `NumHarmonics`: To see how more/less terms affect the approximation
- `Period`: To change the wave frequency
- `samples`: To change the resolution

The visualization is simple but effectively shows how the series approximates a square wave, with more harmonics leading to sharper transitions. For more sophisticated visualization, you could:
- Integrate with a graphics library like `gg`
- Export to a plotting tool

## Configuration Details

- Added `numExamplePoints := 5` as a configuration value alongside `samples := 200`
- Changed the loop to use `numExamplePoints` instead of hardcoded 8
- Set `numExamplePoints` to 5 because with 200 samples, multiplying by 40 gives us valid indices at 0, 40, 80, 120, 160 (5 points total)

### Array Bounds Safety
The program won't panic because:
- With 200 samples, valid indices are 0 to 199
- Using 5 points with step 40 gives us indices: 0, 40, 80, 120, 160
- All these indices are within bounds

### Increasing Example Points
If you want more example points, you could:
- Increase `samples` (e.g., to 320 for 8 points with step 40)
- Decrease the step size (e.g., use i*25 instead of i*40)