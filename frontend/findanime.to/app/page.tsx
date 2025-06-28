"use client";
import React, { useState } from "react";
import { Card, CardBody } from "@heroui/card";
import { Input } from "@heroui/input";

const SearchIcon = () => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="20"
    height="20"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
  >
    <circle cx="11" cy="11" r="8" />
    <line x1="21" y1="21" x2="16.65" y2="16.65" />
  </svg>
);

export default function Home() {
  const [anime, setAnime] = useState("");
  const [recommendations, setRecommendations] = useState([]);

  const handleSearch = async () => {
  if (!anime.trim()) return;
  try {
    const response = await fetch(`http://127.0.0.1:3001/api/v1/public/recommend?anime=${anime}`);
    const data = await response.json();

    // Parse the `message` field if it's a string
    if (data.code === 200 && typeof data.message === "string") {
      const parsed = JSON.parse(data.message);
      setRecommendations(parsed);
    } else {
      setRecommendations([]);
    }
  } catch (error) {
    console.error("Error fetching anime recommendations:", error);
    setRecommendations([]);
  }
};

  return (
    <section className="flex flex-col items-center justify-center gap-6 py-10">
      {/* Heading */}
      <div className="text-center">
        <h1 className="text-5xl font-extrabold mb-2">Find Your&nbsp;
          <span className="text-violet-600">Anime</span>
        </h1>
        <p className="text-gray-400 text-lg">Get instant recommendations with one name</p>
      </div>

      {/* Search Input */}
      <Card className="w-full max-w-4xl shadow-lg rounded-xl p-6 bg-gray-900">
        <CardBody>
          <div className="flex items-center gap-4">
            <Input
              label="Search Anime"
              type="text"
              className="flex-grow"
              value={anime}
              onChange={(e) => setAnime(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && handleSearch()}
            />
            <div
              onClick={handleSearch}
              className="bg-green-500 hover:bg-green-600 p-3 rounded-xl cursor-pointer flex items-center justify-center"
            >
              <SearchIcon />
            </div>
          </div>
        </CardBody>
      </Card>

      {/* Recommendations Grid */}
      {recommendations.length > 0 && (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-6 w-full max-w-6xl mt-6">
          {recommendations.map((anime, index) => (
            <Card key={index} className="bg-gray-800 text-white shadow-md rounded-lg">
              <img
                src={anime.banner}
                alt={anime.title}
                className="rounded-t-lg w-full h-48 object-cover"
              />
              <CardBody>
                <h3 className="text-lg font-semibold mb-1">{anime.title}</h3>
                <p className="text-sm text-gray-300 mb-2 line-clamp-3">{anime.description}</p>
                <p className="text-yellow-400 font-medium">⭐ {anime.rating}</p>
              </CardBody>
            </Card>
          ))}
        </div>
      )}
    </section>
  );
}
