<?php
/**
 *
 */
namespace Nottify;

/**
 * Shared inheritable File handling.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
abstract class File
{
    /**
     * Setup ready for File handling.
     *
     * @param \PDO $pdo
     *  PDO instance for Data working.
     * @param \Nottify\Config $config
     *  Global Config object.
     * @param \getID3 $id3
     *  ID3 meta provider.
     */
    public function __construct(\PDO $pdo, \Nottify\Config $config, \getID3 $id3)
    {
        $this->pdo = $pdo;
        $this->config = $config;
        $this->id3 = $id3;
    }

    /**
     * Load the specific File.
     *
     * @param string $filename
     *  Full path to local file.
     * @return \Nottify\File
     */
    public function loadFromFile($filename)
    {
        $this->filename = $filename;

        $stmt = $this->pdo->prepare('SELECT * FROM nottify WHERE filename=:filename');
        $stmt->execute(array(':filename' => $this->filename));
        $results = $stmt->fetchAll(\PDO::FETCH_ASSOC);

        if (count($results) < 1) {
            // does not exist
        } else {
            $this->database_content = current($results);
        }

        return $this;
    }

    /**
     * Ensure Meta is readily gathered.
     *
     * @throws \RuntimeException
     */
    abstract public function gatherMeta();

    /**
     * Provided and readily available PDO instance.
     *
     * @var \PDO
     */
    protected $pdo;

    /**
     * Shared global Config.
     *
     * @var \Nottify\Config
     */
    protected $config;

    /**
     * Available ID3 object.
     *
     * @var \getID3
     */
    protected $id3;

    /**
     * The current Filename.
     *
     * @var string
     */
    protected $filename = '';

    /**
     * Loaded Database content.
     *
     * @var array
     */
    protected $database_content = array();
}

