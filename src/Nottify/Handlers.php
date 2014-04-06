<?php
/**
 * Provide handling for Urls.
 */
namespace Nottify;

/**
 * Setup handling for Urls within Slim.
 *
 * @package Nottify
 * @author Allan Shone <allan.shone@yahoo.com>
 */
class Handlers
{
    /**
     * Set up for some Handling.
     *
     * @param \Nottify\Config $config
     *  Shared Config object.
     */
    public function __construct(\Nottify\Config $config)
    {
        $this->config = $config;
    }

    /**
     * Set up the handlers.
     *
     * @param \Slim\Slim $app
     *  The current Slim application.
     * @param \Nottify\Engine $nottify
     *  Our Nottify Engine.
     * @return \Slim\Slim
     */
    public function addHandlers(\Slim\Slim $app, \Nottify\Engine $nottify)
    {
        //

        return $app;
    }

    /**
     * Shared Config object.
     *
     * @var \Nottify\Config
     */
    private $config;
}

