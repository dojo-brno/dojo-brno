

def replace(template, mapping):
    if mapping.get(template[1:-1]) == "":
        return ""

    replacement = mapping.get(template[1:-1])
    if replacement:
        return replacement
    return template
