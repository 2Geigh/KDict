#!/bin/bash

# --- 1. Configuration ---
SPEED=0.125            # 0.125 = 800% speed
CONTRAST=1.5           # 1.5 = 50% increase
COLORKEY_SIM=0.9       # Similarity for white transparency
PAUSE_DURATION=0.0     # Seconds to pause on the last frame
TARGET_FPS=30          # Ensures predictable timing and pause length
OUTPUT_DIR="./fast"

mkdir -p "$OUTPUT_DIR"

# --- 2. Filter Component Construction ---
# Normalize visuals first
VISUALS="format=gray,eq=contrast=${CONTRAST},colorkey=white:${COLORKEY_SIM}:0"

# Timing: Force a constant frame rate before reordering
# This makes the "0.5s pause" math exact.
TIMING="fps=${TARGET_FPS},setpts=${SPEED}*PTS"

# Pause Logic: Only add the filter if duration > 0
PAUSE_FILTER=""
if (( $(echo "$PAUSE_DURATION > 0" | bc -l) )); then
    PAUSE_FILTER=",tpad=stop_mode=clone:stop_duration=${PAUSE_DURATION}"
fi

# --- 3. Execution ---
for f in *.gif; do
    echo "Processing: $f..."

    ffmpeg -v error -i "$f" -filter_complex \
        "[0:v]${VISUALS},${TIMING}[v]; \
         [v]split=2[main][first]; \
         [first]select='eq(n,0)'[f0]; \
         [main]select='gt(n,0)'[rest]; \
         [rest][f0]concat=n=2:v=1:a=0${PAUSE_FILTER}[final]; \
         [final]split[inner_a][inner_b]; \
         [inner_a]palettegen=reserve_transparent=1[p]; \
         [inner_b][p]paletteuse" \
        "${OUTPUT_DIR}/fast_${f}"
done

echo "Done. Outputs located in $OUTPUT_DIR"