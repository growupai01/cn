import os
from docx import Document
from docx.document import Document as _Document
from docx.oxml.text.paragraph import CT_P
from docx.oxml.table import CT_Tbl
from docx.table import _Cell, Table
from docx.text.paragraph import Paragraph

DOCX_PATH = os.path.join("sdkdemo", "Warden协议文档.docx")
OUTPUT_PATH = "Warden协议文档.md"

def iter_block_items(parent):
    """
    Generate a reference to each paragraph and table child within parent,
    in document order. Each returned value is an instance of either Table or
    Paragraph.
    """
    if isinstance(parent, _Document):
        parent_elm = parent.element.body
    elif isinstance(parent, _Cell):
        parent_elm = parent._tc
    else:
        raise ValueError("Something's not right")

    for child in parent_elm.iterchildren():
        if isinstance(child, CT_P):
            yield Paragraph(child, parent)
        elif isinstance(child, CT_Tbl):
            yield Table(child, parent)

def format_table(table):
    md_table = []
    # Header
    headers = [cell.text.strip() for cell in table.rows[0].cells]
    md_table.append("| " + " | ".join(headers) + " |")
    md_table.append("| " + " | ".join(["---"] * len(headers)) + " |")
    
    # Rows
    for row in table.rows[1:]:
        row_cells = [cell.text.strip().replace("\n", "<br>") for cell in row.cells]
        md_table.append("| " + " | ".join(row_cells) + " |")
    
    return "\n".join(md_table) + "\n\n"

def convert_docx_to_md(docx_path, md_path):
    if not os.path.exists(docx_path):
        print(f"File not found: {docx_path}")
        return

    doc = Document(docx_path)
    md_content = ""

    print(f"Converting {docx_path} to {md_path}...")

    for block in iter_block_items(doc):
        if isinstance(block, Paragraph):
            text = block.text.strip()
            if not text:
                continue
            
            style_name = block.style.name
            
            if style_name.startswith('Heading 1'):
                md_content += f"# {text}\n\n"
            elif style_name.startswith('Heading 2'):
                md_content += f"## {text}\n\n"
            elif style_name.startswith('Heading 3'):
                md_content += f"### {text}\n\n"
            elif style_name.startswith('Heading 4'):
                md_content += f"#### {text}\n\n"
            elif style_name.startswith('List Paragraph') or style_name.startswith('List Bullet'):
                md_content += f"- {text}\n"
            elif style_name.startswith('List Number'):
                md_content += f"1. {text}\n" # Simplified list handling
            else:
                md_content += f"{text}\n\n"
        
        elif isinstance(block, Table):
            md_content += format_table(block)

    with open(md_path, "w", encoding="utf-8") as f:
        f.write(md_content)
    
    print(f"Successfully created {md_path}")

if __name__ == "__main__":
    convert_docx_to_md(DOCX_PATH, OUTPUT_PATH)
