import os
replacement = "ferru97/libocr"
for dname, dirs, files in os.walk("/home/ferru97/Documents/libocr"):
    for fname in files:
        if fname.endswith(".go") or fname.endswith(".sum") or fname.endswith(".mod"):
            fpath = os.path.join(dname, fname)
            with open(fpath) as f:
                s = f.read()
            s = s.replace("smartcontractkit/libocr", replacement)
            with open(fpath, "w") as f:
                f.write(s)

