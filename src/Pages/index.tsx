import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const Index: React.FC = () => {
    const [selectedImage, setSelectedImage] = useState<File | null>(null);

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedImage(event.target.files[0]);
        }
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        if (selectedImage) {
            console.log('Selected image:', selectedImage.name);

            const formData = new FormData();
            formData.append('image', selectedImage);

            try {
                console.log('Sending request...');
                const response = await fetch('http://localhost:8000/upload', {
                    method: 'POST',
                    body: formData,
                });
                console.log('Response:', response);

            } catch (error) {
                console.log('Error uploading image:', error);
            }
        }
    };

    return (
        <div>
            <div>
                <img id="hero_image" src="/71XtBFClyOL.jpg" alt="Hero" />
            </div>
            <ul id="upload">
                <li>
                    <form onSubmit={handleSubmit}>
                        <input
                            type="file"
                            name="file"
                            onChange={handleFileChange}
                        />
                        <button type="submit">Upload</button>
                    </form>
                </li>
                <li>
                    <Link to="#">View images</Link>
                </li>
            </ul>
        </div>
    );
};

export default Index;
