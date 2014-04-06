<?php
/**
 * Provide MP3 file handling.
 */
namespace Nottify\File;

/**
 * Simple handling of MP3 file type meta detail.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
class Mp3 extends File
{
    /**
     * Use the ID3 to gather Meta information.
     *
     * @return bool
     * @throws \RuntimeException
     */
    public function gatherMeta()
    {
        $analysis = $this->id3->analyze($this->filename);

        if (isset($analysis['error'])) {
            $this->error[] = $analysis['error'];
        }

        $this->track['artist'] = $this->getMetaItem($analysis, 'artist');
        $this->track['title'] = $this->getMetaItem($analysis, 'title');
        $this->track['album'] = $this->getMetaItem($analysis, 'album');

        $this->track['number'] = $this->getMetaItem($analysis, 'track_number');
        if (empty($this->track['number'])) {
            $this->track['number'] = $this->getMetaItem($analysis, 'track');
        }

        $this->track['genre'] = $this->getMetaItem($analysis, 'genre');
        $this->track['file'] = $this->filename;
        $this->track['hash'] = uniqid() . '-' . $this->getMetaItem($analysis, 'md5_data');
        $this->track['playtime'] = $this->getMetaItem($analysis, 'playtime_string');
        $this->track['mime'] = $this->getMetaItem($analysis, 'mime_type');

        if (!empty($this->errors)) {
            throw new \RuntimeException('Errors: ' . implode(', ', $this->errors));
        }

        return true;
    }

    /**
     * Simple ID3 helper for Item retrieval.
     *
     * @param array $id3
     *  Full ID3 info generated.
     * @param string $item
     *  Name of item to retrieve.
     * @return string
     */
    protected function getMetaItem($id3, $item)
    {
        if (isset($id3['tags']['id3v2'][$item])) {
            return $id3['tags']['id3v2'][$item];
        } else if (isset($id3['tags']['id3v1'][$item])) {
            return $id3['tags']['id3v1'][$item];
        } else if (isset($id3['tags']['ape'][$item])) {
            return $id3['tags']['ape'][$item];
        } else if (isset($id3[$item])) {
            return $id3[$item];
        } else {
            $this->errors[] = 'Unable to find ' . $item;

            return '';
        }
    }

    /**
     * The current building track.
     *
     * @var array
     */
    protected $track = array();

    /**
     * Keep track of errors during processing.
     *
     * @var array
     */
    protected $errors = array();
}

