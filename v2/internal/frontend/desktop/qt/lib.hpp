#pragma once

#ifdef __cplusplus
extern "C" {
#endif
extern void appExited(int code);

/* Application */
void *Application_new(char *app_name);
void Application_exec(void *app_ptr);
void Application_quit(void *app_ptr);
char *Application_get_screens(void *app_ptr); // Returns a json blob
/* End Application */

/* Window */
typedef struct Window {
  void *window;
  void *window_layout;
  void *web_engine_view;
} Window;

Window *Window_new(void *app_ptr, char *start_url);
void Window_set_title(void *win_ptr, char *title);
void Window_resize(void *win_ptr, int width, int height);
void Window_set_minimum_size(void *win_ptr, int width, int height);
void Window_hide(void *win_ptr);
void Window_show(void *win_ptr);
void Window_fullscreen(void *win_ptr);
void Window_maximize(void *win_ptr);
void Window_close(void *win_ptr);
/* End Window */

/* WebEngineView */
void WebEngineView_load_url(void *web_engine_ptr, char *url);
void WebEngineView_reload(void *web_engine_ptr);
void WebEngineView_run_js(void *web_engine_ptr, char *script);
/* End WebEngineView */

void fix_signal(int signum);
void install_signal_handlers();

void bye(void* ptr);

#ifdef __cplusplus
}
#endif