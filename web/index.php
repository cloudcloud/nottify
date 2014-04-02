<?php
/**
 * Nottify media player.
 *
 * @author Allan Shone <allan.shone@yahoo.com>
 */

$autoload = realpath(__DIR__ . '/../vendor/autoload.php');
if (!file_exists($autoload)) {
    // well, composer hasn't been loaded
    header('HTTP/1.1 404 Not Found');
    exit('Nottify has not been configured or installed correctly.');
}

require_once $autoload;

// make sure we have a config
$config_file = realpath(__DIR__ . '/../config.ini');
try {
    $config = new \Nottify\Config($config_file);
} catch (\Exception $e) {
    header('HTTP/1.1 500 Internal Server Error');
    exit($e->getMessage());
}

// provide some basic IP filtering
if (!$config->isValidUser($_SERVER)) {
    header('HTTP/1.0 403 Forbidden');
    exit('Access Forbidden');
}

// fixing a bug with the server
$_SERVER['SCRIPT_NAME'] = '/index.php';

// set up the app server
$app = new \Slim\Slim(array(
    'debug' => $config->getBool('main.debug'),
    'mode' => $config->determineMode($config->getBool('main.debug')),
    'cookies.encrypt' => true,
    'cookies.httponly' => true,
    'cookies.secret_key' => $config->getString('secret.cookie'),
));
$app->setName('nottify');

// set up Nottify
$nottify = new \Nottify\Engine($config);

// provide handlers
$handlers = new \Nottify\Handlers($config);
$app = $handlers->addHandlers($app, $nottify);
$app->run();

