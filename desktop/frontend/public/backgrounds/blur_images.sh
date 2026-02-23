#!/bin/bash

# Configuration
OUTPUT_DIR="./blurs"
BLUR_RATIO=0.005
EXTENSIONS=("jpg" "jpeg" "png" "bmp" "webp")

mkdir -p "$OUTPUT_DIR"

# Processing
shopt -s nullglob
for ext in "${EXTENSIONS[@]}"; do
    for img in *."$ext"; do
        # 1. Use 'w' and 'h' constants (iw/ih are not valid for boxblur).
        # 2. Escape the comma in the min() function to prevent filter-chain splitting.
        # 3. Increase luma_power to 4 for a smoother, high-fidelity Gaussian approximation.
        
        ffmpeg -v error -i "$img" -vf \
        "format=rgb24,boxblur=luma_radius=min(w\,h)*$BLUR_RATIO:luma_power=4,format=yuv420p" \
        "$OUTPUT_DIR/${img%.*}_blur.${img##*.}" -y
    done
done
