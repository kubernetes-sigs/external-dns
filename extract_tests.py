
import re

filename = "provider/cloudflare/cloudflare_test.go.bak"
target_funcs = [
    "TestBuildCustomHostnameNewParams",
    "TestSubmitCustomHostnameChanges",
    "TestGroupByNameAndTypeWithCustomHostnames_MX",
    "TestProviderPropertiesIdempotency"
]

with open(filename, "r") as f:
    content = f.read()

# Naive extraction assuming func starts at newline and invalid Go code doesn't exist (brace matching)
# We track brace depth.

extracted = ""

lines = content.splitlines()
i = 0
while i < len(lines):
    line = lines[i]
    match = re.match(r"^func (Test[a-zA-Z0-9_]+)\(", line)
    if match:
        func_name = match.group(1)
        if func_name in target_funcs:
            # Start extracting
            start_line = i
            brace_count = 0
            # Scan inside function
            # We need to account for the first line having braces
            # A simple way for Go functions usually formatted: ends with ^}$
            
            # Let's count braces to be robust
            func_content = []
            
            while i < len(lines):
                line = lines[i]
                func_content.append(line)
                brace_count += line.count("{")
                brace_count -= line.count("}")
                
                if brace_count == 0:
                    # Function ended
                    extracted += "\n" + "\n".join(func_content) + "\n"
                    break
                i += 1
    i += 1

print(extracted)
