// Unused usings removed
using System.IO;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.EntityFrameworkCore;
using HeroesApi.Models;
using Microsoft.Extensions.FileProviders;

using System.Text;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using Microsoft.IdentityModel.Logging;
using Microsoft.AspNetCore.Authentication;
using System.Linq;


namespace HeroesApi
{
    public class Startup
    {
        //readonly string MyAllowSpecificOrigins = "_myAllowSpecificOrigins";
        private const string Name = "HeroesCORS";
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddCors(options =>
            {
                options.AddPolicy(name: Name, //MyAllowSpecificOrigins,
                    builder =>
                    {
                        // AllowAnyMethod needed for PUT request when changing hero name
                        builder.WithOrigins("http://localhost:80")
                            .AllowAnyHeader()
                            .AllowAnyMethod()
                            .AllowCredentials();
                    });
            });

            services.AddDbContext<HeroContext>(opt =>
               opt.UseInMemoryDatabase("HeroesList"));

            var appSettingsSection = Configuration.GetSection("AppSettings");
            services.Configure<AppSettings>(appSettingsSection);

            IdentityModelEventSource.ShowPII = true;

            // configure jwt authentication
            var appSettings = appSettingsSection.Get<AppSettings>();
            var key = Encoding.ASCII.GetBytes(appSettings.SecretKey);
            var audience = appSettings.Audience;
            services.AddAuthentication(this.SetAuthenticationOptions).AddJwtBearer(jwtBearerOptions => this.SetJwtBearerOptions(jwtBearerOptions, key, audience));

            services.AddControllers();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();
            app.UseAuthentication();
	        app.UseFileServer(new FileServerOptions
            {
                FileProvider = new PhysicalFileProvider(
                  Path.Combine(Directory.GetCurrentDirectory(), "AngularApp")),
                RequestPath = "",

            });		
            //app.UseHttpsRedirection();
            DefaultFilesOptions options = new DefaultFilesOptions();
            options.DefaultFileNames.Clear();
            options.DefaultFileNames.Add("index.html");
            app.UseDefaultFiles(options);
            app.UseStaticFiles();  
            app.UseCors(Name); //MyAllowSpecificOrigins); 
            app.UseAuthorization();
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
                // catch non existing path (except files, e.g. ".../dashboard.png" - this still gets the error 404)
                endpoints.MapFallbackToFile("/index.html");
                // ... and catch file errors too    
                endpoints.MapFallbackToFile("{*path:file}", "/index.html");
            });
            
        }
            private void SetAuthenticationOptions(AuthenticationOptions authenticationOptions)
        {
            authenticationOptions.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
            authenticationOptions.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
        }

        private void SetJwtBearerOptions(JwtBearerOptions jwtBearerOptions, byte[] key, string audience)
        {
            jwtBearerOptions.RequireHttpsMetadata = false;
            jwtBearerOptions.SaveToken = true;
            jwtBearerOptions.TokenValidationParameters = new TokenValidationParameters
            {
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(key),
                ValidateIssuer = false,
                ValidateAudience = true,

                AudienceValidator = (audiences, token, validationParameters) =>
                {
                    return audiences.Contains(audience);
                },
            };
        }

        private class AppSettings
        {
            public string SecretKey { get; set; }
            public string Audience { get; set; }
        }
    }
}

