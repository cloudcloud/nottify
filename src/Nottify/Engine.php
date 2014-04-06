<?php
/**
 * Provide the driving force.
 */
namespace Nottify;

/**
 * Full handling for all web detail.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
class Engine
{
    /**
     * Simple Engine starter.
     *
     * @param \Nottify\Config $config
     *  Global Config object.
     */
    public function __construct(\Nottify\Config $config)
    {
        $this->config = $config;
    }

    /**
     * Provide a fresh Track.
     *
     * @return \Nottify\Track
     */
    public function getTrack()
    {
        $track = new \Nottify\Track($this->config);

        return $track;
    }

    /**
     * Shared Config object.
     *
     * @var \Nottify\Config
     */
    protected $config;
}

