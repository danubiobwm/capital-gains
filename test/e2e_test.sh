        #!/usr/bin/env bash
        set -euo pipefail

        INPUT_FILE=${1:-test/examples_input.txt}
        EXPECTED_FILE=${2:-test/expected_output.txt}
        BIN=${3:-./capital-gains}

        if [ ! -f "$BIN" ]; then
          echo "Binary $BIN not found. Build first (make build)."
          exit 2
        fi

        if [ ! -f "$INPUT_FILE" ]; then
          echo "Input file $INPUT_FILE not found"
          exit 2
        fi

        if [ ! -f "$EXPECTED_FILE" ]; then
          echo "Expected output file $EXPECTED_FILE not found"
          exit 2
        fi

        # Run the program and capture output lines
        mapfile -t got_lines < <("$BIN" < "$INPUT_FILE")

        mapfile -t exp_lines < "$EXPECTED_FILE"

        if [ "${#got_lines[@]}" -ne "${#exp_lines[@]}" ]; then
          echo "Line count mismatch: got ${#got_lines[@]} expected ${#exp_lines[@]}"
          echo "---- GOT ----"
          printf '%s
' "${got_lines[@]}"
          echo "---- EXPECTED ----"
          printf '%s
' "${exp_lines[@]}"
          exit 1
        fi

        for i in "${!exp_lines[@]}"; do
          # trim spaces
          got=$(echo "${got_lines[$i]}" | tr -d '[:space:]')
          exp=$(echo "${exp_lines[$i]}" | tr -d '[:space:]')
          if [ "$got" != "$exp" ]; then
            echo "Mismatch at line $((i+1)):"
            echo " GOT: ${got_lines[$i]}"
            echo "EXP: ${exp_lines[$i]}"
            exit 1
          fi
        done

        echo "E2E tests passed!"
