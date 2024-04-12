

from svglib.svglib import svg2rlg
from reportlab.graphics import renderPM
from PIL import Image
import os
import glob

def convert_svg_to_png(input_directory, output_directory):
    os.makedirs(output_directory, exist_ok=True)  # Ensure output directory exists
    svg_files = glob.glob(os.path.join(input_directory, '*.svg'))  # List all SVG files

    for svg_file in svg_files:
        # Convert SVG to an intermediate format (ReportLab Drawing)
        drawing = svg2rlg(svg_file)
        file_name = os.path.basename(svg_file)
        png_path = os.path.join(output_directory, file_name.replace('.svg', '.png'))

        # Render Drawing to a PNG file using Pillow
        renderPM.drawToFile(drawing, png_path, fmt='PNG')

    print("Conversion complete!")

    # Specify the directory containing the SVG files
input_directory = 'avatar'
output_directory = 'avatar-png'




convert_svg_to_png(input_directory, output_directory)
