#!/bin/bash

# Loop semua folder di direktori saat ini
for dir in */; do
    # Cek apakah folder berisi Makefile
    if [ -f "$dir/Makefile" ]; then
        echo "🔧 Running 'make build' in $dir ..."
        (
            cd "$dir" || exit
            if make build; then
                echo "✅ Success: $dir"
            else
                echo "❌ Failed: $dir (skipping...)"
            fi
        )
    else
        echo "⚠️ Skip $dir (no Makefile found)"
    fi
done
