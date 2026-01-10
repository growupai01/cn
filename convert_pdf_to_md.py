import google.generativeai as genai
import os
from pypdf import PdfReader

# Configuration
API_KEY = "sk-9dcab62bb8014686a1b026e7e8939321"
API_ENDPOINT = 'http://127.0.0.1:8045'
# Trying gemini-3-pro-high as requested, but keeping fallback option in mind
MODEL_NAME = "gemini-3-pro-high" 
PDF_PATH = os.path.join("sdkdemo", "Warden协议文档.pdf")
OUTPUT_PATH = "Warden协议文档.md"

def extract_text_from_pdf(pdf_path):
    print(f"Reading from: {os.path.abspath(pdf_path)}")
    try:
        reader = PdfReader(pdf_path)
        text = ""
        for page in reader.pages:
            text += page.extract_text() + "\n"
        return text
    except Exception as e:
        print(f"Error reading PDF: {e}")
        return None

def convert_to_markdown(text):
    # Unset proxies
    os.environ.pop("HTTP_PROXY", None)
    os.environ.pop("HTTPS_PROXY", None)
    os.environ.pop("http_proxy", None)
    os.environ.pop("https_proxy", None)

    try:
        genai.configure(
            api_key=API_KEY,
            transport='rest',
            client_options={'api_endpoint': API_ENDPOINT}
        )
        
        model = genai.GenerativeModel(MODEL_NAME)
        
        prompt = (
            "You are a helpful assistant that converts document text into Markdown.\n"
            "Please convert the following text from a protocol document into a well-formatted Markdown file.\n"
            "Ensure that:\n"
            "1. Headers are correctly formatted (H1, H2, etc.).\n"
            "2. Lists and tables are preserved.\n"
            "3. Content is identical to the source.\n"
            "4. No introductory or concluding remarks, just the markdown content.\n\n"
            f"Source Text:\n{text[:50000]}"
        )
        
        response = model.generate_content(prompt)
        return response.text
    except Exception as e:
        import traceback
        traceback.print_exc()
        print(f"Error calling API: {e}")
        return None

def main():
    if not os.path.exists(PDF_PATH):
        print(f"File not found: {PDF_PATH}")
        return

    print(f"Extracting text from {PDF_PATH}...")
    text = extract_text_from_pdf(PDF_PATH)
    if not text:
        print("No text extracted.")
        return

    print(f"Extracted {len(text)} characters. Converting to Markdown using Gemini...")
    md_content = convert_to_markdown(text)
    
    if md_content:
        # Clean up code blocks if the model wrapped it
        if md_content.startswith("```markdown"):
            md_content = md_content.replace("```markdown", "", 1)
        if md_content.startswith("```"):
             md_content = md_content.replace("```", "", 1)
        if md_content.endswith("```"):
            md_content = md_content[:-3]

        with open(OUTPUT_PATH, "w", encoding="utf-8") as f:
            f.write(md_content.strip())
        print(f"Successfully saved to {os.path.abspath(OUTPUT_PATH)}")
    else:
        print("Conversion failed.")

if __name__ == "__main__":
    main()
