#!/bin/bash

for dir in */; do
    [ -d "$dir" ] || continue

    if [ -f "${dir}Makefile" ]; then
        echo "🔧 Running 'make build' in '$dir' ..."
        (
            cd "$dir" || { echo "❌ Failed to enter directory: $dir"; echo; exit 1; }
            if make build; then
                echo "✅ Success: $dir"
                echo
            else
                echo "❌ Failed: $dir (skipping...)"
                echo
            fi
        )
    else
        echo "⚠️ Skipping '$dir' (no Makefile found)"
        echo
    fi
done

echo "🏁 All done!"
