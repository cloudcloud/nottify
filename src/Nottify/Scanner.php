<?php
/**
 * Provide Scan detail.
 */
namespace Nottify;

/**
 * Provide full scanning capabilities across the file system.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
class Scanner
{
    /**
     * Provide ease of beginning Scan processing.
     *
     * @param \PDO $pdo
     *  Provided PDO instance.
     * @param \Nottify\Config $config
     *  Readily available global Config object.
     * @param \getID3 $id3
     *  Object for acquiring full ID3 information.
     */
    public function __construct(\PDO $pdo, \Nottify\Config $config, \getID3 $id3)
    {
        $this->pdo = $pdo;
        $this->config = $config;
        $this->id3 = $id3;

        //
    }

    /**
     * Complete a full file system trawl for specific files.
     *
     * @param string $location
     *  Base file system path to look for files.
     * @param string $file_type
     *  Specific file type to look for.
     * @return array
     */
    public function searchFileSystem($location, $file_type)
    {
        $file_listing = explode("\n", `find {$directory} -name "*.{$file_type}"`);

        $files = array();
        foreach ($file_listing as $file) {
            $class_name = "\\Nottify\\File\\" . ucfirst($file_type);
            $f = new $class_name($this->pdo, $this->config, $this->id3);
            $files[] = $f->loadFromFile($file);
        }

        return $files;
    }

    /**
     * Complete a processing of provided file list.
     *
     * @param array $files
     *  An array containing File objects.
     * @return integer
     */
    public function processFiles($files)
    {
        foreach ($files as $file) {
            $file->gatherMeta();
        }
    }

    /**
     * Shared PDO instance.
     *
     * @var \PDO
     */
    private $pdo;

    /**
     * Global Config instance.
     *
     * @var \Nottify\Config
     */
    private $config;

    /**
     * ID3 processor.
     *
     * @var \getID3
     */
    private $id3;
}

