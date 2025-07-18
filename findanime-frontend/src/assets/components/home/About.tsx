import React from 'react';

interface AboutProps {
  positionStyle?: React.CSSProperties;
  maxWidth?: string | number; // optional
}

export default function About({ positionStyle = {} }: AboutProps) {
  return (
    <div
  className="w-full max-w-5xl py-12 text-white font-sans space-y-6"
  style={{
    position: 'absolute',
    top: positionStyle.top || 'auto',
    bottom: positionStyle.bottom || 'auto',
    left: positionStyle.left || 'auto',
    right: positionStyle.right || 'auto',
    margin: '0 auto',
    ...positionStyle,
  }}
>

      <h1 className="text-2xl md:text-3xl font-bold">
        Findanime.to – The best site to watch anime online for Free
      </h1>

      <p className="text-white/80">
        Do you know that according to Google, the monthly search volume for anime related topics is up to over 1 Billion times? Anime is famous worldwide and it is no wonder we’ve seen a sharp rise in the number of free anime streaming sites.
      </p>

      <p className="text-white/80">
        Just like free online movie streaming sites, anime watching sites are not created equally, some are better than the rest, so we’ve decided to build Findanime.to to be one of the best free anime streaming site for all anime fans on the world.
      </p>

      <div>
        <h2 className="text-xl font-bold mt-8">1/ What is Findanime.to?</h2>
        <p className="text-white/80 mt-2">
          Findanime.to is a free site to watch anime and you can even download subbed or dubbed anime in ultra HD quality without any registration or payment. By having only one ads in all kinds, we are trying to make it the safest site for free anime.
        </p>
      </div>

      <div>
        <h2 className="text-xl font-bold mt-8">2/ Is Findanime.to safe?</h2>
        <p className="text-white/80 mt-2">
          Yes we are, we do have only one Ads to cover the server cost and we keep scanning the ads 24/7 to make sure all are clean, If you find any ads that is suspicious, please forward us the info and we will remove it.
        </p>
      </div>

      <div>
        <h2 className="text-xl font-bold mt-8">
          3/ So what make Findanime.to the best site to watch anime free online?
        </h2>

        <p className="text-white/80 mt-2">
          Before building Findanime.to, we’ve checked many other free anime sites, and learnt from them. We only keep the good things and remove all the bad things from all the competitors, to put it in our Findanime website. Let’s see how we’re so confident about being the best site for anime streaming:
        </p>

        <ul className="list-disc pl-6 space-y-2 mt-4 text-white/80">
          <li><strong className="text-white">Safety:</strong> We try our best to not having harmful ads on Findanime.</li>
          <li>
            <strong className="text-white">Content library:</strong> Our main focus is anime. You can find here popular, classic, as well as current titles from all genres such as action, drama, kids, fantasy, horror, mystery, police, romance, school, comedy, music, game and many more. All these titles come with English subtitles or are dubbed in many languages.
          </li>
          <li>
            <strong className="text-white">Quality/Resolution:</strong> All titles are in excellent resolution, the best quality possible. Findanime.to also has a quality setting function to make sure our users can enjoy streaming no matter how fast your Internet speed is. You can stream the anime at 360p if your Internet is being ridiculous, Or if it is good, you can go with 720p or even 1080p anime.
          </li>
          <li>
            <strong className="text-white">Streaming experience:</strong> Compared to other anime streaming sites, the loading speed at Findanime.to is faster. Downloading is just as easy as streaming, you won’t have any problem saving the videos to watch offline later.
          </li>
          <li>
            <strong className="text-white">Updates:</strong> We updates new titles as well as fulfill the requests on a daily basis so be warned, you will never run out of what to watch on Findanime.
          </li>
          <li>
            <strong className="text-white">User interface:</strong> Our UI and UX makes it easy for anyone, no matter how old you are, how long have you been on the Internet. Literally, you can figure out how to navigate our site after a quick look. If you want to watch a specific title, search for it via the search box. If you want to look for suggestions, you can use the site’s categories or simply scroll down for new releases.
          </li>
          <li>
            <strong className="text-white">Device compatibility:</strong> Findanime works right on both your mobile and desktop. However, we’d recommend you use your desktop for a smoother streaming experience.
          </li>
          <li>
            <strong className="text-white">Customer care:</strong> We are in active mode 24/7. You can always contact us for any help, query, or business-related inquiry. On our previous projects, we were known for our great customer service as we were quick to fix broken links or upload requested content.
          </li>
        </ul>
      </div>

      <p className="mt-6 text-white/80">
        So if you’re looking for a trustworthy and safe site for your Anime streaming, let’s give Findanime.to a try. And if you like us, please help us to spread the words and do not forget to bookmark our site.
      </p>

      <p className="text-white/80">Thank you!</p>
    </div>
  );
}
