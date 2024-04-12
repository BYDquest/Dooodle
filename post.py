# pip install Mastodon.py
# pip install python-dotenv

from mastodon import Mastodon
import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Replace these variables with your own details
api_base_url = 'https://mastodon.social'  # your Mastodon instance URL

# Retrieve environment variables
client_id = os.getenv('CLIENT_ID')
client_secret = os.getenv('CLIENT_SECRET')
access_token = os.getenv('ACCESS_TOKEN')

# Authenticate the Mastodon client
mastodon = Mastodon(
    client_id=client_id,
    client_secret=client_secret,
    access_token=access_token,
    api_base_url=api_base_url
)

# Path to your image
image_path = './combined-image/2.png'  # Make sure to update this path
# Upload an image
media = mastodon.media_post(image_path)
# Create a post with the uploaded image
mastodon.status_post('', media_ids=[media])

print("Posted to Mastodon successfully!")
