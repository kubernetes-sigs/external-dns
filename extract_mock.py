
import re

filename = "provider/cloudflare/cloudflare_test.go.bak"

with open(filename, "r") as f:
    lines = f.readlines()

extracted = []
capturing = False
brace_depth = 0

# We want to capture:
# 1. type mockCloudFlareClient struct
# 2. func NewMockCloudFlareClient...
# 3. func (m *mockCloudFlareClient) ...

mock_struct_pattern = re.compile(r"^type mockCloudFlareClient struct")
mock_func_pattern = re.compile(r"^func \(m \*mockCloudFlareClient\)")
factory_pattern = re.compile(r"^func NewMockCloudFlareClient")

i = 0
while i < len(lines):
    line = lines[i]
    if mock_struct_pattern.match(line) or mock_func_pattern.match(line) or factory_pattern.match(line):
        capturing = True
    
    if capturing:
        extracted.append(line)
        brace_depth += line.count("{")
        brace_depth -= line.count("}")
        if brace_depth == 0 and line.strip() != "":
             # Check if we assume valid Go (end of struct/func is })
             if line.strip() == "}" or line.strip() == "},":
                 capturing = False
                 extracted.append("\n") # Spacer
    i += 1

# Also filter out cloudflarev0 usages in the extracted mock and replace with CustomHostname if needed?
# My previous analysis showed .bak used cloudflarev0.CustomHostname in some places but maybe it was aliased?
# Let's just output it first.

print("".join(extracted))
